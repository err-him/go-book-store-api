package models

import (
	"book-store-api/api/utils"
	"context"
	"time"
)

type Books struct {
	Id          int64                  `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	ISBN        *string                `json:"isbn,omitempty"`
	Language    *string                `json:"langauge,omitempty"`
	PublisherId *uint16                `json:"publisher_id,omitempty"`
	PublishedAt *utils.MysqlFormatDate `json:"published_at,omitempty"`
	BookGenre   *BookGenre             `json:"book_genre,omitempty"`
	BookAuthor  *BookAuthor            `json:"book_author,omitempty"`
	Prices      *BookPrice             `json:"prices,omitempty"`
	Other       *BookOthers            `json:"other,omitempty"`
	Status      *bool                  `json:"status,omitempty"`
	CreatedAt   *time.Time             `json:"created_at,omitempty"`
	UpdatedAt   *time.Time             `json:"updated_at,omitempty"`
}

type BookGenre struct {
	Id *[]int `json:"id,omitempty"`
}

type BookPrice struct {
	NewPrice *float32 `json:"new,omitempty"`
	OldPrice *float32 `json:"old,omitempty"`
}

type BookOthers struct {
	Quantity    *uint8  `json:"quantity,omitempty"`
	Type        *string `json:"type,omitempty"`
	NumberPages *uint8  `json:"number_pages,omitempty"`
}

type BookAuthor struct {
	Id *[]int `json:"id,omitempty"`
}

type BooksRepo interface {
	Add(ctx context.Context, r *Books) (*Books, error)
}
