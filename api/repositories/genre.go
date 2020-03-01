package repositories

import (
	"book-store-api/api/models"
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"
)

type GenreRepo struct {
	db *sql.DB
}

func NewGenreRepo(db *sql.DB) models.GenreRepo {
	return &GenreRepo{
		db: db,
	}
}

func (g *GenreRepo) Create(ctx context.Context, r *models.CreateGenre) (*models.CreateGenreResponse, error) {
	slug, err := g.generateGenreSlug(r.Name)
	if err != nil {
		return nil, err
	}
	//insert record in database
	res, err := g.db.Exec(`Insert into genre (name, slug) value (?, ?)`, strings.TrimSpace(r.Name), slug)
	if err != nil {
		return nil, err
	}
	insertedId, _ := res.LastInsertId()
	payload := &models.CreateGenreResponse{
		Id:   insertedId,
		Slug: slug,
	}
	return payload, nil
}

/**
 * [func description]
 * @param  {[type]} g [description]
 * @return {[type]}   [description]
 */
func (g *GenreRepo) Update(ctx context.Context, r *models.Genre) (bool, error) {
	//update record in database
	rows, err := g.db.Exec(`update genre set name=?, status = ?, updated_at = ? where id=?`, strings.TrimSpace(*r.Name), r.Status, time.Now(), r.Id)
	if err != nil {
		return false, err
	}
	row, _ := rows.RowsAffected()
	if row > 0 {
		return true, nil
	}
	return false, models.ErrNotFound
}

/**
 * [func description]
 * @param  {[type]} g [description]
 * @return {[type]}   [description]
 */
func (g *GenreRepo) GetAll(ctx context.Context, limit int64, offset int64) ([]*models.Genre, error) {
	rows, err := g.db.Query(`select * from genre where status = 1 LIMIT  ?,?`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.Genre, 0)
	for rows.Next() {
		data := new(models.Genre)
		err := rows.Scan(
			&data.Id,
			&data.Name,
			&data.Slug,
			&data.Status,
			&data.CreatedAt,
			&data.UpdatedAt,
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
 * @param  {[type]} g [description]
 * @return {[type]}   [description]
 */
func (g *GenreRepo) GetOne(ctx context.Context, id string) (*models.Genre, error) {
	data := &models.Genre{}
	genreId, _ := strconv.ParseInt(id, 10, 64)
	row := g.db.QueryRow(`select * from genre where id = ? OR slug = ? `, genreId, id)
	err := row.Scan(
		&data.Id,
		&data.Name,
		&data.Slug,
		&data.Status,
		&data.CreatedAt,
		&data.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return data, nil
		}
		return nil, err
	}
	return data, nil
}

/**
 * [func description]
 * @param  {[type]} g [description]
 * @return {[type]}   [description]
 */
func (g *GenreRepo) Delete(ctx context.Context, id string) (bool, error) {
	genreId, _ := strconv.ParseInt(id, 10, 64)
	res, err := g.db.Exec(`update genre set status = 0 where id=? OR slug = ?`, genreId, id)
	rows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if rows >= 0 {
		return true, nil
	}
	return false, nil
}

/**
 * [func description]
 * @param  {[type]} g [description]
 * @return {[type]}   [description]
 */
func (g *GenreRepo) generateGenreSlug(slug string) (string, error) {

	var id int32
	slug = strings.TrimSpace(slug)
	slug = strings.ToLower(slug)
	//change all spaces to - in order to create a valid Slug
	slug = strings.Replace(slug, " ", "-", -1)
	//check if slug already present in DB
	row := g.db.QueryRow(`SELECT id FROM genre WHERE slug=?`, slug)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return slug, nil
		} else {
			return "", err
		}
	}
	return "", models.ErrAlreadyPresent
}
