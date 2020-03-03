package repositories

import (
	cons "book-store-api/api/constants"
	"book-store-api/api/handler"
	"book-store-api/api/models"
	"context"
	"database/sql"
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
		res, err := tx.Exec(`Insert into books (name, isbn, price, langauge, quantity, book_type, old_price,publisher_id,number_pages,published_at,image) value (?, ?, ?, ?, ?, ? , ?, ?, ?, ?, ?)`, strings.TrimSpace(*r.Name), *r.ISBN, *r.Prices.NewPrice, *r.Language, *r.Other.Quantity, *r.Other.Type, *r.Prices.OldPrice, *r.PublisherId, *r.Other.NumberPages, r.PublishedAt.Format(cons.MYSQL_DATE_FORMAT), strings.TrimSpace(*r.Image))
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
		res, err := tx.Exec(`update  books set name = ?, isbn = ?, price = ?, langauge = ?, quantity = ?, old_price = ?,publisher_id = ?,number_pages = ?,published_at = ? , status = ?, book_type = ?, updated_at = ?, image = ? where id = ?`, strings.TrimSpace(*r.Name), *r.ISBN, *r.Prices.NewPrice, *r.Language, *r.Other.Quantity, *r.Prices.OldPrice, *r.PublisherId, *r.Other.NumberPages, r.PublishedAt.Format(cons.MYSQL_DATE_FORMAT), *r.Status, *r.Other.Type, time.Now(), strings.TrimSpace(*r.Image), r.Id)
		if err != nil {
			return err
		}
		row, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if row <= 0 {
			return models.ErrNotFound
		}

		//delete all the rows first and then reinsert again
		res, err = tx.Exec("delete from book_genre where book_id = ?", resId)
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
 * @param  {[type]} b [description]
 * @return {[type]}   [description]
 */
func (b *BookRepo) GetBookDetailById(ctx context.Context, id int64) ([]*models.Books, error) {
	payload := make([]*models.Books, 0)
	data := new(models.Books)
	data.BookPublishers = new(models.BookPublisher)
	data.Prices = new(models.BookPrice)
	data.Other = new(models.BookOthers)
	data.BookGenreRes = make([]*models.BookMetaRes, 0)
	data.BookAuthorRes = make([]*models.BookMetaRes, 0)
	row := b.db.QueryRow(`SELECT b.*, p.name as publication_name FROM books b inner join publications as p ON p.id = b.publisher_id where b.id = ?`, id)

	err := row.Scan(
		&data.Id,
		&data.Name,
		&data.ISBN,
		&data.Prices.NewPrice,
		&data.Language,
		&data.Other.Quantity,
		&data.Prices.OldPrice,
		&data.Other.Type,
		&data.BookPublishers.Id,
		&data.Image,
		&data.Status,
		&data.Other.NumberPages,
		&data.PublishedAt,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.BookPublishers.Name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return payload, nil
		}
		return nil, err
	}
	data.BookGenreRes, err = b.getBookGenreDetails(id)
	if err != nil {
		return nil, err
	}
	data.BookAuthorRes, err = b.getBookAuthorDetails(id)
	if err != nil {
		return nil, err
	}
	payload = append(payload, data)
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

/**
 * [func description]
 * @param  {[type]} g [description]
 * @return {[type]}   [description]
 */
func (b *BookRepo) GetAll(ctx context.Context, limit int64, offset int64) ([]*models.Books, error) {
	rows, err := b.db.Query(`SELECT b.*, p.name as publication_name FROM books b inner join publications as p ON p.id = b.publisher_id LIMIT  ?,?`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.Books, 0)
	for rows.Next() {
		data := new(models.Books)
		data.BookPublishers = new(models.BookPublisher)
		data.Prices = new(models.BookPrice)
		data.Other = new(models.BookOthers)
		data.BookGenreRes = make([]*models.BookMetaRes, 0)
		data.BookAuthorRes = make([]*models.BookMetaRes, 0)
		err := rows.Scan(
			&data.Id,
			&data.Name,
			&data.ISBN,
			&data.Prices.NewPrice,
			&data.Language,
			&data.Other.Quantity,
			&data.Prices.OldPrice,
			&data.Other.Type,
			&data.BookPublishers.Id,
			&data.Image,
			&data.Status,
			&data.Other.NumberPages,
			&data.PublishedAt,
			&data.CreatedAt,
			&data.UpdatedAt,
			&data.BookPublishers.Name,
		)
		if err != nil {
			return nil, err
		}
		data.BookGenreRes, err = b.getBookGenreDetails(data.Id)
		if err != nil {
			return nil, err
		}
		data.BookAuthorRes, err = b.getBookAuthorDetails(data.Id)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (b *BookRepo) SearchBookByName(ctx context.Context, q string) ([]*models.Books, error) {

	rows, err := b.db.Query(`SELECT id, name, image FROM books where name LIKE ?`, "%"+q+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.Books, 0)
	for rows.Next() {
		data := new(models.Books)
		err := rows.Scan(
			&data.Id,
			&data.Name,
			&data.Image,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

/**
 * [func description]
 * @param  {[type]} b [description]
 * @return {[type]}   [description]
 */
func (b *BookRepo) getBookGenreDetails(bookId int64) ([]*models.BookMetaRes, error) {
	rows, err := b.db.Query(`SELECT g.id, g.name FROM genre g inner join book_genre as bg on g.id = bg.genre_id where bg.book_id = ?`, bookId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.BookMetaRes, 0)
	for rows.Next() {
		data := new(models.BookMetaRes)
		err := rows.Scan(
			&data.Id,
			&data.Name,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

/**
 * [func description]
 * @param  {[type]} b [description]
 * @return {[type]}   [description]
 */
func (b *BookRepo) getBookAuthorDetails(bookId int64) ([]*models.BookMetaRes, error) {
	rows, err := b.db.Query(`SELECT a.id, a.name FROM author a inner join book_author as ba on a.id = ba.author_id where ba.book_id = ?`, bookId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.BookMetaRes, 0)
	for rows.Next() {
		data := new(models.BookMetaRes)
		err := rows.Scan(
			&data.Id,
			&data.Name,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}
