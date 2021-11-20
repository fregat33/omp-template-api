package repo

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/est-rent-api/internal/model"
)

var (
	//ErrNotFound is error when estate rent not found
	ErrNotFound = errors.New("estate rent not found")
)

// Repo is DAO for Template
type Repo interface {
	CreateRent(ctx context.Context, rent *model.Rent) (*model.Rent, error)
	DescribeRent(ctx context.Context, templateID uint64) (*model.Rent, error)
	ListRent(ctx context.Context, fromRentID, limit uint64) ([]model.Rent, error)
	RemoveRent(ctx context.Context, rentID uint64) (bool, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

//DescribeRent get rent
func (r *repo) DescribeRent(ctx context.Context, rentID uint64) (*model.Rent, error) {
	return nil, nil
}

//CreateRent add new rent
func (r *repo) CreateRent(ctx context.Context, rent *model.Rent) (*model.Rent, error) {
	return nil, nil
}

//RemoveRent remove existed rent
func (r *repo) RemoveRent(ctx context.Context, rentID uint64) (bool, error) {
	return false, nil
}

//ListRent get list of rents
func (r *repo) ListRent(ctx context.Context, fromRentID, limit uint64) ([]model.Rent, error) {
	return nil, nil
}


