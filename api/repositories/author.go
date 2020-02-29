package repositories

import (
	"book-store-api/api/models"
	"database/sql"
)

type AuthorRepo struct {
	db *sql.DB
}

func NewAuthorRepo(db *sql.DB) models.AuthorRepo {
	return &AuthorRepo{
		db: db,
	}
}
