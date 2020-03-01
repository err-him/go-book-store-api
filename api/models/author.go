package models

import (
	"book-store-api/api/utils"
	"context"
	"time"
)

type AuthorCreateReq struct {
	Id        *int64     `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Slug      *string    `json:"slug,omitempty"`
	About     About      `json:"about,omitempty"`
	Status    *bool      `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type About struct {
	YearsActive *utils.MysqlFormatDate `json:"years_active,omitempty"`
	Language    *string                `json:"language"`
	Personal    Personal               `json:"personal"`
}

type Personal struct {
	Dob  *utils.MysqlFormatDate `json:"dob,omitempty"`
	Info *string                `json:"info"`
}

type CreateAuthorRes struct {
	Id   int64  `json:"id"`
	Slug string `json:"slug"`
}

type AuthorRepo interface {
	Create(ctx context.Context, r *AuthorCreateReq) (*CreateAuthorRes, error)
	GetAll(ctx context.Context, limit int64, offset int64) ([]*AuthorCreateReq, error)
	GetOne(ctx context.Context, id string) ([]*AuthorCreateReq, error)
	Delete(ctx context.Context, id string) (bool, error)
	Update(ctx context.Context, r *AuthorCreateReq) (bool, error)
}
