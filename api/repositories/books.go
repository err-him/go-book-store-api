package repositories

import (
	"book-store-api/api/models"
	"database/sql"
)

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) models.BooksRepo {
	return &BookRepo{
		db: db,
	}
}
