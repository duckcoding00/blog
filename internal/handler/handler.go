package handler

import (
	"net/http"

	"github.com/duckcoding00/blog/internal/service"
)

type Handler struct {
	Blog interface {
		Create(http.ResponseWriter, *http.Request)
		Update(http.ResponseWriter, *http.Request)
		GetBlog(http.ResponseWriter, *http.Request)
		GetBlogs(http.ResponseWriter, *http.Request)
		Delete(http.ResponseWriter, *http.Request)
	}
}

func NewHandler() Handler {
	service := service.NewService()
	return Handler{
		Blog: &BlogHandler{
			service: service,
		},
	}
}
