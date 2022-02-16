package usecases

import (
	"context"

	"github.com/ochom/jumia-interview-task/database"
	"github.com/ochom/jumia-interview-task/models/dto"
	"github.com/ochom/jumia-interview-task/utils"
)

// CustomersUsecase ...
type CustomersUsecase interface {
	GetNumbers(ctx context.Context, code, state string) (resp []*dto.FormattedNumber, err error)
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

func (uc *impl) GetNumbers(ctx context.Context, code, state string) (resp []*dto.FormattedNumber, err error) {
	customers, err := uc.repo.GetAllNumbers(ctx)
	if err != nil {
		return nil, err
	}

	// if country code not all and not empty
	if code != "all" && code != "" {
		customers = utils.FilterInCountry(code, customers)
	}

	formattedNumbers := utils.FormatNumbers(customers)

	// get numbers based on state
	if state == "all" || state == "" {
		resp = formattedNumbers
		return
	}

	for _, v := range formattedNumbers {
		formtedNumber := *v
		if formtedNumber.State == state {
			resp = append(resp, &formtedNumber)
		}
	}
	return
}
