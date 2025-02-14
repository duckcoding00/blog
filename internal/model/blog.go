package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

type BlogPayload struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

func (u BlogPayload) Validate() error {
	return Validate.Struct(u)
}

type BlogUpdatePayload struct {
	Title *string `json:"title" validate:"omitempty"`
	Body  *string `json:"body" validate:"omitempty"`
}

func (u BlogUpdatePayload) Validate() error {
	return Validate.Struct(u)
}

type BlogResponse struct {
	ID        int32     `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
