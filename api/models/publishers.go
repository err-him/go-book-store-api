package models

import (
	"book-store-api/api/utils"
	"context"
	"time"
)

//CreatePubsRes  create publishers request response
type CreatePubsRes struct {
	Id   int64  `json:"name"`
	Slug string `json:"slug"`
}
type UpdateReqPublish struct {
	Id     *int64          `json:"id"`
	Name   *string         `json:"name,omitempty"`
	Meta   *CreateMetaData `json:"meta,omitempty"`
	Status *bool           `json:"status,omitempty"`
}
type GetPubsResponse struct {
	Id        int64          `json:"id,omitempty"`
	Name      string         `json:"name,omitempty"`
	Meta      MetaGetReponse `json:"meta,omitempty"`
	Status    bool           `json:"status,omitempty"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
}

type MetaGetReponse struct {
	Desc         string `json:"desc,omitempty"`
	Slug         string `json:"slug,omitempty"`
	FoundingDate string `json:"founding_date,omitempty"`
}

//ReqPublish req publish request body
type ReqPublish struct {
	Name   *string         `json:"name,omitempty"`
	Meta   *CreateMetaData `json:"meta,omitempty"`
	Status *bool           `json:"status,omitempty"`
}

//CreateMetaData create meta data
type CreateMetaData struct {
	FoundingDate *utils.MysqlFormatDate `json:"founding_date,omitempty"`
	Desc         *string                `json:"desc,omitempty"`
}

type PubsRepo interface {
	Create(ctx context.Context, r *ReqPublish) (*CreatePubsRes, error)
	GetAll(ctx context.Context, limit int64, offset int64) ([]*GetPubsResponse, error)
	GetOne(ctx context.Context, id string) ([]*GetPubsResponse, error)
	Delete(ctx context.Context, id string) (bool, error)
	Update(ctx context.Context, r *UpdateReqPublish) (bool, error)
}
