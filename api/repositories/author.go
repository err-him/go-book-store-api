package repositories

import (
	cons "book-store-api/api/constants"
	"book-store-api/api/models"
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"
)

type AuthorRepo struct {
	db *sql.DB
}

func NewAuthorRepo(db *sql.DB) models.AuthorRepo {
	return &AuthorRepo{
		db: db,
	}
}

func (a *AuthorRepo) Create(ctx context.Context, r *models.AuthorCreateReq) (*models.CreateAuthorRes, error) {
	slug, err := a.generateSlug(*r.Name)
	if err != nil {
		return nil, err
	}

	//insert record in database
	res, err := a.db.Exec(`Insert into author (name, years_active, slug, dob, about,language) value (?, ?, ?, ?, ?, ?)`, strings.TrimSpace(*r.Name), r.About.YearsActive.Format(cons.MYSQL_DATE_FORMAT), slug, r.About.Personal.Dob.Format(cons.MYSQL_DATE_FORMAT), strings.TrimSpace(*r.About.Personal.Info), *r.About.Language)

	if err != nil {
		return nil, err
	}
	insertedId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	payload := &models.CreateAuthorRes{
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
func (a *AuthorRepo) GetAll(ctx context.Context, limit int64, offset int64) ([]*models.AuthorCreateReq, error) {
	rows, err := a.db.Query(`select * from author where status = 1 LIMIT  ?,?`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.AuthorCreateReq, 0)
	for rows.Next() {
		data := new(models.AuthorCreateReq)
		err := rows.Scan(
			&data.Id,
			&data.Name,
			&data.About.YearsActive,
			&data.Slug,
			&data.About.Personal.Dob,
			&data.About.Personal.Info,
			&data.About.Language,
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
func (a *AuthorRepo) GetOne(ctx context.Context, id string) ([]*models.AuthorCreateReq, error) {
	payload := make([]*models.AuthorCreateReq, 0)
	data := new(models.AuthorCreateReq)
	authorId, _ := strconv.ParseInt(id, 10, 64)
	row := a.db.QueryRow(`select * from author where id = ? OR slug = ? `, authorId, id)
	err := row.Scan(
		&data.Id,
		&data.Name,
		&data.About.YearsActive,
		&data.Slug,
		&data.About.Personal.Dob,
		&data.About.Personal.Info,
		&data.About.Language,
		&data.Status,
		&data.CreatedAt,
		&data.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return payload, nil
		}
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
func (a *AuthorRepo) Delete(ctx context.Context, id string) (bool, error) {
	pId, _ := strconv.ParseInt(id, 10, 64)
	res, err := a.db.Exec(`update author set status = 0 where id=? OR slug = ?`, pId, id)
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
 * @param  {[type]} p [description]
 * @return {[type]}   [description]
 */

func (a *AuthorRepo) Update(ctx context.Context, r *models.AuthorCreateReq) (bool, error) {
	//insert record in database
	res, err := a.db.Exec(`Update author set name = ?, years_active = ?, dob = ?, about = ?,language = ?, status = ?, updated_at = ? where id = ?`, strings.TrimSpace(*r.Name), r.About.YearsActive.Format(cons.MYSQL_DATE_FORMAT), r.About.Personal.Dob.Format(cons.MYSQL_DATE_FORMAT), strings.TrimSpace(*r.About.Personal.Info), *r.About.Language, *r.Status, time.Now(), *r.Id)
	if err != nil {
		return false, err
	}
	row, _ := res.RowsAffected()
	if row > 0 {
		return true, nil
	}
	return false, models.ErrNotFound
}

/**
 * [func description]
 * @param  {[type]} a [description]
 * @return {[type]}   [description]
 */
func (a *AuthorRepo) generateSlug(slug string) (string, error) {

	var id int32
	slug = strings.TrimSpace(slug)
	slug = strings.ToLower(slug)
	//change all spaces to - in order to create a valid Slug
	slug = strings.Replace(slug, " ", "-", -1)
	//check if slug already present in DB
	row := a.db.QueryRow(`SELECT id FROM author WHERE slug=?`, slug)
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
