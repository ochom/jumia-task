package usecases

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/ochom/jumia-interview-task/database"
	"github.com/ochom/jumia-interview-task/database/mock"
	"github.com/ochom/jumia-interview-task/models"
	"github.com/ochom/jumia-interview-task/models/dto"
)

var repo database.Repository
var fakeRepo mock.FakeRepository

func Test_impl_GetAll(t *testing.T) {
	repo = &fakeRepo

	uc := New(repo)
	ctx := context.Background()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []dto.FormattedNumber
		wantErr bool
	}{
		{
			name:    "unable to get all customers",
			args:    args{ctx: ctx},
			want:    nil,
			wantErr: true,
		},
		{
			name: "happy got all customers",
			args: args{ctx: ctx},
			want: []dto.FormattedNumber{
				{Country: "Mozambique", Code: "+258", State: "OK", Phone: "847651504"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if tt.name == "unable to get all customers" {
			fakeRepo.GetAllNumbersFn = func(ctx context.Context) ([]*models.Customer, error) {
				return nil, fmt.Errorf("unable to get phone number")
			}
		}

		if tt.name == "happy got all customers" {
			fakeRepo.GetAllNumbersFn = func(ctx context.Context) ([]*models.Customer, error) {
				return []*models.Customer{
					{Name: "Test", Phone: "(258) 847651504"},
				}, nil
			}
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := uc.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_GetByCountry(t *testing.T) {
	repo = &fakeRepo

	uc := New(repo)
	ctx := context.Background()

	type args struct {
		ctx  context.Context
		code string
	}
	tests := []struct {
		name    string
		args    args
		want    []dto.FormattedNumber
		wantErr bool
	}{{
		name:    "unable to get all customers",
		args:    args{ctx: ctx, code: "254"},
		want:    nil,
		wantErr: true,
	},
		{
			name: "happy got all customers",
			args: args{ctx: ctx, code: "258"},
			want: []dto.FormattedNumber{
				{Country: "Mozambique", Code: "+258", State: "OK", Phone: "847651504"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if tt.name == "unable to get all customers" {
			fakeRepo.GetAllNumbersFn = func(ctx context.Context) ([]*models.Customer, error) {
				return nil, fmt.Errorf("unable to get phone number")
			}
		}

		if tt.name == "happy got all customers" {
			fakeRepo.GetAllNumbersFn = func(ctx context.Context) ([]*models.Customer, error) {
				return []*models.Customer{
					{Name: "Test", Phone: "(258) 847651504"},
				}, nil
			}
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := uc.GetByCountry(tt.args.ctx, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetByCountry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetByCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}
