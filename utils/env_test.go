package utils

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {

	os.Setenv("ping", "pong")

	type args struct {
		name        string
		alternative string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy env set",
			args: args{
				name:        "ping",
				alternative: "pang",
			},
			want: "pong",
		},
		{
			name: "sad env not set",
			args: args{
				name:        "pong",
				alternative: "pang",
			},
			want: "pang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnv(tt.args.name, tt.args.alternative); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
