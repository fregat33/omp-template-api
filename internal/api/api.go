package api

import (
	"context"
	"errors"
	"github.com/ozonmp/est-rent-api/internal/model"
	"github.com/shopspring/decimal"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozonmp/est-rent-api/internal/repo"

	pb "github.com/ozonmp/est-rent-api/pkg/est-rent-api"
)

var (
	totalTemplateNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "est_rent_api_rent_not_found_total",
		Help: "Total number of templates that were not found",
	})
)

type rentAPI struct {
	pb.UnimplementedEstRentApiServiceServer
	repo repo.Repo
}

// NewTemplateAPI returns api of est-rent-api service
func NewTemplateAPI(r repo.Repo) pb.EstRentApiServiceServer {
	return &rentAPI{repo: r}
}

func (o *rentAPI) CreateRentV1(
	ctx context.Context,
	req *pb.CreateRentV1Request,
) (*pb.CreateRentV1Response, error) {
	log.Debug().Msg("CreateRentV1 called")

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	price, _ := decimal.NewFromString(req.PriceDecimal)
	rent := &model.Rent{
		RenterId:   req.RenterId,
		ObjectInfo: req.ObjectInfo,
		ObjectType: model.RentObjectType(req.ObjectType),
		Period:     time.Duration(req.Period),
		Price:      price,
	}

	rent, err := o.repo.CreateRent(ctx, rent)
	if err != nil {
		log.Error().Err(err).Msg("Create")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("CreateRentV1 finished")

	return &pb.CreateRentV1Response{
		Data: rentToPb(rent),
	}, nil
}

func (o *rentAPI) DescribeRentV1(
	ctx context.Context,
	req *pb.DescribeRentV1Request,
) (*pb.DescribeRentV1Response, error) {
	log.Debug().Msg("DescribeRentV1 called")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeRentV1 invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	rent, err := o.repo.DescribeRent(ctx, req.RentId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeRentV1: error on repo.DescribeRent")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if rent == nil {
		log.Debug().Uint64("templateId", req.RentId).Msg("rent not found")
		totalTemplateNotFound.Inc()

		return nil, status.Error(codes.NotFound, "rent not found")
	}

	log.Debug().Msg("DescribeRentV1 finished")

	return &pb.DescribeRentV1Response{Data: rentToPb(rent)}, nil
}

func (o *rentAPI) ListRentV1(
	ctx context.Context,
	req *pb.ListRentV1Request,
) (*pb.ListRentV1Response, error) {
	log.Debug().Msg("ListRentV1: called")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("ListRentV1: error validate")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	rents, err := o.repo.ListRent(ctx, req.FromRentId, req.Limit)
	if err != nil {
		log.Error().Err(err).Msg("ListRentV1: error repo.ListRent")
		return nil, status.Error(codes.Internal, err.Error())
	}

	list := make([]*pb.Rent, len(rents))
	for i := range rents {
		list[i] = rentToPb(&rents[i])
	}

	log.Debug().Msg("ListRentV1: finished")

	return &pb.ListRentV1Response{
		Items: list,
	}, nil
}

func (o *rentAPI) RemoveRentV1(
	ctx context.Context,
	req *pb.RemoveRentV1Request,
) (*pb.RemoveRentV1Response, error) {
	log.Debug().Msg("RemoveRentV1 called")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemoveRentV1: error validation")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	success, err := o.repo.RemoveRent(ctx, req.RentId)
	if err != nil {
		log.Error().Err(err).Msg("RemoveRentV1: error repo.RemoveRent")

		if errors.Is(err, repo.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Bool("deleted", success).Msgf("RemoveRentV1: finished")

	return &pb.RemoveRentV1Response{}, nil
}

func rentToPb(rent *model.Rent) *pb.Rent {
	return &pb.Rent{
		Id:           rent.ID,
		RenterId:     rent.RenterId,
		ObjectInfo:   rent.ObjectInfo,
		ObjectType:   rent.ObjectType.String(),
		Period:       int64(rent.Period),
		PriceDecimal: rent.Price.String(),
	}
}
