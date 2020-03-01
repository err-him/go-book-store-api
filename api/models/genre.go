package models

import (
	"context"
	"time"
)

type CreateGenre struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CreateGenreResponse struct {
	Id   int64  `json:"id,omitempty"`
	Slug string `json:"slug,omitempty"`
}

//Genre Details
type Genre struct {
	Id        *int64     `json:"id,string,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Slug      *string    `json:"slug,omitempty"`
	Status    *bool      `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type GenreRepo interface {
	Create(ctx context.Context, g *CreateGenre) (*CreateGenreResponse, error)
	Update(ctx context.Context, g *Genre) (bool, error)
	GetAll(ctx context.Context, limit int64, offset int64) ([]*Genre, error)
	GetOne(ctx context.Context, id string) (*Genre, error)
	Delete(ctx context.Context, id string) (bool, error)
}
