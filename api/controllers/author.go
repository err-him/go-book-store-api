package controllers

import (
	"book-store-api/api/models"
	r "book-store-api/api/repositories"
	"book-store-api/config/driver"
)

type Author struct {
	authorRepo models.AuthorRepo
}

func NewAuthorHandler(db *driver.DB) *Author {
	return &Author{
		authorRepo: r.NewAuthorRepo(db.SQL),
	}
}
