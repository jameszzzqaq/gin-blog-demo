package service

import (
	"testing"

	"github.com/yu1er/gin-blog/model"
)

func TestCheckAuthExist(t *testing.T) {
	type args struct {
		a *model.Auth
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"存在", args{&model.Auth{Username: "test", Password: "test123456"}}, true},
		{"不存在", args{&model.Auth{Username: "test", Password: "no"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckAuthExist(tt.args.a)
			if got != tt.want {
				t.Errorf("CheckAuthExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
