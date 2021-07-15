package service

import (
	"reflect"
	"testing"

	"github.com/yu1er/gin-blog/model"
	"github.com/yu1er/gin-blog/model/request"
)

func TestGetTagsPage(t *testing.T) {
	type args struct {
		info request.TagListGet
	}
	tests := []struct {
		name  string
		args  args
		want  []model.Tag
		want1 int
	}{
		// TODO: Add test cases.
		{"test", args{request.TagListGet{request.Page{0, 2}, ""}}, nil, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, _ := GetTagsPage(tt.args.info)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTagsPage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetTagsPage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
