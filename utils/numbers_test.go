package utils

import (
	"testing"
)

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
