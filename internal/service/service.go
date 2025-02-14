package service

import (
	"context"
	"path/filepath"
	"runtime"

	"github.com/duckcoding00/blog/internal/model"
)

type Service struct {
	Blog interface {
		Create(context.Context, model.BlogPayload) (model.BlogResponse, error)
		Update(context.Context, int32, model.BlogUpdatePayload) (model.BlogResponse, error)
		GetBlog(context.Context, int32) (model.BlogResponse, error)
		GetBlogs(context.Context) ([]model.BlogResponse, error)
		Delete(context.Context, int32) error
	}
}

func NewService() Service {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	dataPath := filepath.Join(filepath.Dir(currentDir), "data", "blog.json")
	return Service{
		Blog: &BlogService{
			FilePath: dataPath,
		},
	}
}
