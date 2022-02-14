package database

import (
	"context"

	"github.com/ochom/jumia-interview-task/models"
	"gorm.io/gorm"
)

// Repository interface for database related methods
type Repository interface {
	GetAllNumbers(ctx context.Context) ([]*models.Customer, error)
}

type impl struct {
	db *gorm.DB
}

// New creates a new instance of the repository
func New(db *gorm.DB) Repository {
	return &impl{
		db: db,
	}
}

func (r *impl) GetAllNumbers(ctx context.Context) ([]*models.Customer, error) {
	var customers []*models.Customer

	err := r.db.Raw("SELECT id, name, phone FROM customer").Scan(&customers).Error

	if err != nil {
		return nil, err
	}

	return customers, nil
}
