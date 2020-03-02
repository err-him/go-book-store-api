package repositories

import (
	cons "book-store-api/api/constants"
	"book-store-api/api/handler"
	"book-store-api/api/models"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) models.BooksRepo {
	return &BookRepo{
		db: db,
	}
}

/**
 * [func description]
 * @param  {[type]} b [description]
 * @return {[type]}   [description]
 */
func (b *BookRepo) Add(ctx context.Context, r *models.Books) (*models.Books, error) {
	var resId int64
	err := handler.WithTransaction(b.db, func(tx handler.Transaction) error {
		//insert record into `books` tables
		res, err := tx.Exec(`Insert into books (name, isbn, price, langauge, quantity, book_type, old_price,publisher_id,number_pages,published_at) value (?, ?, ?, ?, ?, ? , ?, ?, ?, ?)`, strings.TrimSpace(*r.Name), *r.ISBN, *r.Prices.NewPrice, *r.Language, *r.Other.Quantity, *r.Other.Type, *r.Prices.OldPrice, *r.PublisherId, *r.Other.NumberPages, r.PublishedAt.Format(cons.MYSQL_DATE_FORMAT))
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

/**
 * [func description]
 * @param  {[type]} b [description]
 * @return {[type]}   [description]
 */
func (b *BookRepo) Update(ctx context.Context, r *models.Books) (*models.Books, error) {
	var resId = r.Id
	err := handler.WithTransaction(b.db, func(tx handler.Transaction) error {
		//insert record into `books` tables
		res, err := tx.Exec(`update  books set name = ?, isbn = ?, price = ?, langauge = ?, quantity = ?, old_price = ?,publisher_id = ?,number_pages = ?,published_at = ? , status = ?, book_type = ?, updated_at = ? where id = ?`, strings.TrimSpace(*r.Name), *r.ISBN, *r.Prices.NewPrice, *r.Language, *r.Other.Quantity, *r.Prices.OldPrice, *r.PublisherId, *r.Other.NumberPages, r.PublishedAt.Format(cons.MYSQL_DATE_FORMAT), *r.Status, *r.Other.Type, time.Now(), r.Id)
		if err != nil {
			return err
		}
		row, err := res.RowsAffected()
		if err != nil {
			return err
		}
		fmt.Println("row", row)
		if row <= 0 {
			return models.ErrNotFound
		}

		//delete all the rows first and then reinsert again
		res, err = tx.Exec("delete from book_genre where book_id = ?", resId)
		if err != nil {
			fmt.Println("Err", err)
			return err
		}
		//insert into `book_genre`
		for _, v := range *r.BookGenre.Id {

			res, err = tx.Exec("INSERT INTO book_genre (book_id, genre_id) VALUES(?, ?)", resId, v)
			if err != nil {
				return err
			}
		}
		//delete all the rows first and then reinsert again
		res, err = tx.Exec("delete from book_author where book_id = ?", resId)
		if err != nil {
			return err
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

/**
 * [func description]
 * @param  {[type]} g [description]
 * @return {[type]}   [description]
 */
func (b *BookRepo) Delete(ctx context.Context, id int64) (bool, error) {
	res, err := b.db.Exec(`update books set status = 0 where id=?`, id)
	row, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if row <= 0 {
		return false, models.ErrNotFound
	}
	return true, nil
}
