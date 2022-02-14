package utils

import (
	"reflect"
	"testing"

	"github.com/ochom/jumia-interview-task/models"
	"github.com/ochom/jumia-interview-task/models/dto"
)

func TestFormatNumber(t *testing.T) {
	customer1 := models.Customer{
		ID:    1,
		Phone: "(212) 654642448",
		Name:  "Walid Hammadi",
	}

	customer2 := models.Customer{
		ID:    2,
		Phone: "(2120) 698054317",
		Name:  "Yosaf Karrouch",
	}

	type args struct {
		customer models.Customer
	}
	tests := []struct {
		name string
		args args
		want *dto.FormattedNumber
	}{
		{
			name: "sad regex not match",
			args: args{
				customer2,
			},
			want: nil,
		},
		{
			name: "happy got formatted",
			args: args{
				customer1,
			},
			want: &dto.FormattedNumber{
				Country: "Moroco",
				Phone:   "654642448",
				Code:    "+212",
				State:   "OK",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatNumber(tt.args.customer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCountry(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
	}{
		{
			name:    "got country",
			args:    args{"(212) 654642448"},
			wantNil: false,
		},
		{
			name:    "error in number country",
			args:    args{"21(2 6007989253"},
			wantNil: true,
		},
		{
			name:    "no country match",
			args:    args{"(2120) 698054317"},
			wantNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCountry(tt.args.phone)
			if (got == nil) != tt.wantNil {
				t.Errorf("GetCountry() = %v, wantNil %v", got, tt.wantNil)
			}
		})
	}
}
