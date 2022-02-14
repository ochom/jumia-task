package mock

import (
	"context"

	"github.com/ochom/jumia-interview-task/models"
)

// FakeRepository mocked repository
type FakeRepository struct {
	GetAllNumbersFn func(ctx context.Context) ([]*models.Customer, error)
}

// GetAllNumbers ...
func (f *FakeRepository) GetAllNumbers(ctx context.Context) ([]*models.Customer, error) {
	return f.GetAllNumbersFn(ctx)
}
