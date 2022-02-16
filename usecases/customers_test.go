package usecases

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/ochom/jumia-interview-task/database/mock"
	"github.com/ochom/jumia-interview-task/models"
	"github.com/ochom/jumia-interview-task/models/dto"
)

var fakeRepo mock.FakeRepository

func Test_impl_GetNumbers(t *testing.T) {
	repo := &fakeRepo

	uc := New(repo)

	correctResp1 := []*dto.FormattedNumber{
		{
			Country: "Moroco",
			State:   "NOK",
			Code:    "+212",
			Phone:   "6007989253",
		},
	}

	correctResp2 := []*dto.FormattedNumber{
		{
			Country: "Mozambique",
			State:   "OK",
			Code:    "+258",
			Phone:   "847651504",
		},
		{
			Country: "Mozambique",
			State:   "NOK",
			Code:    "+258",
			Phone:   "042423566",
		},
	}

	correctResp3 := []*dto.FormattedNumber{
		{
			Country: "Mozambique",
			State:   "OK",
			Code:    "+258",
			Phone:   "847651504",
		},
	}

	type args struct {
		ctx   context.Context
		code  string
		state string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []*dto.FormattedNumber
		wantErr  bool
	}{
		{
			name: "could not get all numbers",
			args: args{
				ctx:   context.Background(),
				code:  "all",
				state: "all",
			},
			wantResp: nil,
			wantErr:  true,
		},
		{
			name: "should get all numbers",
			args: args{
				ctx:   context.Background(),
				code:  "all",
				state: "all",
			},
			wantResp: correctResp1,
			wantErr:  false,
		},
		{
			name: "should get Mozambique numbers",
			args: args{
				ctx:   context.Background(),
				code:  "258",
				state: "all",
			},
			wantResp: correctResp2,
			wantErr:  false,
		},
		{
			name: "should get OK Mozambique numbers",
			args: args{
				ctx:   context.Background(),
				code:  "258",
				state: "OK",
			},
			wantResp: correctResp3,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		if tt.name == "could not get all numbers" {
			fakeRepo.GetAllNumbersFn = func(ctx context.Context) ([]*models.Customer, error) {
				return nil, fmt.Errorf("database error, unable to get numbers")
			}
		}

		if tt.name == "should get all numbers" {
			fakeRepo.GetAllNumbersFn = func(ctx context.Context) ([]*models.Customer, error) {
				customers := []*models.Customer{
					{
						Name:  "Test",
						Phone: "(212) 6007989253",
					},
				}
				return customers, nil
			}
		}

		if tt.name == "should get Mozambique numbers" {
			fakeRepo.GetAllNumbersFn = func(ctx context.Context) ([]*models.Customer, error) {
				customers := []*models.Customer{
					{
						Name:  "Test",
						Phone: "(212) 6007989253",
					},
					{
						Name:  "Test 2",
						Phone: "(258) 847651504",
					},
					{
						Name:  "Test 2",
						Phone: "(258) 042423566",
					},
				}
				return customers, nil
			}
		}

		if tt.name == "should get OK Mozambique numbers" {
			fakeRepo.GetAllNumbersFn = func(ctx context.Context) ([]*models.Customer, error) {
				customers := []*models.Customer{
					{
						Name:  "Test",
						Phone: "(212) 6007989253",
					},
					{
						Name:  "Test 2",
						Phone: "(258) 847651504",
					},
					{
						Name:  "Test 2",
						Phone: "(258) 042423566",
					},
				}
				return customers, nil
			}
		}

		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := uc.GetNumbers(tt.args.ctx, tt.args.code, tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("impl.GetNumbers() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
