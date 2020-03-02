package repositories

import (
	cons "book-store-api/api/constants"
	"book-store-api/api/handler"
	"book-store-api/api/models"
	"context"
	"database/sql"
	"strings"
)

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) models.BooksRepo {
	return &BookRepo{
		db: db,
	}
}

func (b *BookRepo) Add(ctx context.Context, r *models.Books) (*models.Books, error) {
	var resId int64
	err := handler.WithTransaction(b.db, func(tx handler.Transaction) error {
		//insert record into `books` tables
		res, err := tx.Exec(`Insert into books (name, isbn, price, langauge, quantity, old_price,publisher_id,number_pages,published_at) value (?, ?, ?, ?, ?, ?, ?, ?, ?)`, strings.TrimSpace(*r.Name), *r.ISBN, *r.Prices.NewPrice, *r.Language, *r.Other.Quantity, *r.Prices.OldPrice, *r.PublisherId, *r.Other.NumberPages, r.PublishedAt.Format(cons.MYSQL_DATE_FORMAT))
		if err != nil {
			return err
		}
		resId, err = res.LastInsertId()
		if err != nil {
			return err
		}
		//insert into `book_genre`
		for _, v := range *r.BookGenre.Id {
			res, err = tx.Exec("INSERT INTO book_genre (book_id, genre_id) VALUES(?, ?)", resId, v)
			if err != nil {
				return err
			}
		}

		//insert into `book_author`
		for _, v := range *r.BookAuthor.Id {
			res, err = tx.Exec("INSERT INTO book_author (book_id, author_id) VALUES(?, ?)", resId, v)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		handler.HandleError(err)
		return nil, err
	}
	payload := &models.Books{
		Id: resId,
	}
	return payload, nil
}
