package usecases

import (
	"context"

	"github.com/ochom/jumia-interview-task/database"
	"github.com/ochom/jumia-interview-task/models/dto"
	"github.com/ochom/jumia-interview-task/utils"
)

// CustomersUsecase ...
type CustomersUsecase interface {
	GetAll(ctx context.Context) ([]dto.FormattedNumber, error)
	GetByCountry(ctx context.Context, code string) ([]dto.FormattedNumber, error)
}

type impl struct {
	repo database.Repository
}

// New creates a new instance of customer usecase
func New(repo database.Repository) CustomersUsecase {
	return &impl{
		repo: repo,
	}
}

func (uc *impl) GetAll(ctx context.Context) ([]dto.FormattedNumber, error) {
	customers, err := uc.repo.GetAllNumbers(ctx)
	if err != nil {
		return nil, err
	}

	formattedNumbers := []dto.FormattedNumber{}

	for _, customer := range customers {
		formattedNumber := utils.FormatNumber(*customer)
		if formattedNumber != nil {
			formattedNumbers = append(formattedNumbers, *formattedNumber)
		}
	}

	return formattedNumbers, nil
}

func (uc *impl) GetByCountry(ctx context.Context, code string) ([]dto.FormattedNumber, error) {
	validNumbers, err := uc.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	inCountry := []dto.FormattedNumber{}
	for _, v := range validNumbers {
		if v.Code == "+"+code {
			inCountry = append(inCountry, v)
		}
	}

	return inCountry, nil
}
