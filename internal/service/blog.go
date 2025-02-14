package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/duckcoding00/blog/internal/model"
)

type BlogService struct {
	FilePath string
}

func (s *BlogService) Create(ctx context.Context, payload model.BlogPayload) (model.BlogResponse, error) {
	blogs, err := s.GetBlogs(ctx)
	if err != nil {
		return model.BlogResponse{}, err
	}

	var id int32 = 1
	if len(blogs) > 0 {
		id = blogs[len(blogs)-1].ID + 1
	}

	blog := model.BlogResponse{
		ID:        id,
		Title:     payload.Title,
		Body:      payload.Body,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	blogs = append(blogs, blog)

	file, err := os.Create(s.FilePath)
	if err != nil {
		return model.BlogResponse{}, err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "   ")
	if err := encoder.Encode(blogs); err != nil {
		return model.BlogResponse{}, err
	}

	return blog, nil
}

func (s *BlogService) Update(ctx context.Context, id int32, payload model.BlogUpdatePayload) (model.BlogResponse, error) {
	blogs, err := s.GetBlogs(ctx)
	if err != nil {
		return model.BlogResponse{}, err
	}

	for i, blog := range blogs {
		if blog.ID == id {

			if payload.Title != nil {
				blogs[i].Title = *payload.Title
			}

			if payload.Body != nil {
				blogs[i].Body = *payload.Body
			}

			blogs[i].UpdatedAt = time.Now().UTC()

			file, err := os.Create(s.FilePath)
			if err != nil {
				return model.BlogResponse{}, err
			}
			defer file.Close()

			if err := json.NewEncoder(file).Encode(blogs); err != nil {
				return model.BlogResponse{}, err
			}

			return blogs[i], nil
		}
	}

	return model.BlogResponse{}, errors.New("blog not found")
}

func (s *BlogService) GetBlog(ctx context.Context, id int32) (model.BlogResponse, error) {
	blogs, err := s.GetBlogs(ctx)
	if err != nil {
		return model.BlogResponse{}, err
	}

	for _, blog := range blogs {
		if blog.ID == id {
			return blog, nil
		}
	}

	return model.BlogResponse{}, errors.New("blog not found")
}

func (s *BlogService) GetBlogs(ctx context.Context) ([]model.BlogResponse, error) {
	if _, err := os.Stat(s.FilePath); os.IsNotExist(err) {
		file, err := os.Create(s.FilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create file :%v", err)
		}

		defer file.Close()

		if _, err := file.Write([]byte("[]")); err != nil {
			return nil, fmt.Errorf("failed to write initial data: %v", err)
		}

		return []model.BlogResponse{}, nil
	}

	files, err := os.Open(s.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, err
		}
		return nil, err
	}
	defer files.Close()

	fileInfo, err := files.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %v", err)
	}

	if fileInfo.Size() == 0 {
		writeFile, err := os.Create(s.FilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create file : %v", err)
		}
		defer writeFile.Close()

		if _, err := writeFile.Write([]byte("[]")); err != nil {
			return nil, fmt.Errorf("failed to write initial data:%v", err)
		}

		return []model.BlogResponse{}, nil
	}
	var blogs []model.BlogResponse
	if err := json.NewDecoder(files).Decode(&blogs); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (s *BlogService) Delete(ctx context.Context, id int32) error {
	blogs, err := s.GetBlogs(ctx)
	if err != nil {
		return fmt.Errorf("failed to get blogs: %v", err)
	}

	found := false
	var filteredBlog []model.BlogResponse
	for _, blog := range blogs {
		if blog.ID != id {
			filteredBlog = append(filteredBlog, blog)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("blog with id %d not found", id)
	}

	file, err := os.Create(s.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "   ")
	if err := encoder.Encode(filteredBlog); err != nil {
		return err
	}

	return nil
}
