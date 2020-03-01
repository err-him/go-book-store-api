package controllers

import (
	"book-store-api/api/models"
	r "book-store-api/api/repositories"
	"book-store-api/config/driver"
)

type Books struct {
	bookRepo models.BooksRepo
}

func NewBookHandler(db *driver.DB) *Books {
	return &Books{
		bookRepo: r.NewBookRepo(db.SQL),
	}
}
