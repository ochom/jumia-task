package database

import (
	"context"
	"reflect"
	"testing"

	"github.com/ochom/jumia-interview-task/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// connect to test database
func initRepo() Repository {
	db, err := gorm.Open(sqlite.Open("../sample.db"), &gorm.Config{})
	utils.FailOnError(err)

	r := New(db)
	return r
}

func Test_impl_GetAllNumbers(t *testing.T) {

	r := initRepo()

	ctx := context.Background()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "happy got customers",
			args:    args{ctx: ctx},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllNumbers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetAllNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got) > 0, tt.want) {
				t.Errorf("impl.GetAllNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
