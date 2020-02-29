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

type PubsRepo struct {
	db *sql.DB
}

func NewPubsRepo(db *sql.DB) models.PubsRepo {
	return &PubsRepo{
		db: db,
	}
}

func (p *PubsRepo) Create(ctx context.Context, r *models.ReqPublish) (*models.CreatePubsRes, error) {
	slug, err := p.generatePubsSlug(*r.Name)
	if err != nil {
		return nil, err
	}

	//insert record in database
	res, err := p.db.Exec(`Insert into publications (name, slug, founding_date, description) value (?, ?, ?, ?)`, strings.TrimSpace(*r.Name), slug, r.Meta.FoundingDate.Format(cons.MYSQL_DATE_FORMAT), *r.Meta.Desc)

	if err != nil {
		return nil, err
	}
	insertedId, _ := res.LastInsertId()
	payload := &models.CreatePubsRes{
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
func (p *PubsRepo) GetAll(ctx context.Context, limit int64, offset int64) ([]*models.GetPubsResponse, error) {
	rows, err := p.db.Query(`select * from publications where status = 1 LIMIT  ?,?`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.GetPubsResponse, 0)
	for rows.Next() {
		data := new(models.GetPubsResponse)
		err := rows.Scan(
			&data.Id,
			&data.Name,
			&data.Meta.Slug,
			&data.Meta.FoundingDate,
			&data.Meta.Desc,
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
func (p *PubsRepo) GetOne(ctx context.Context, id string) ([]*models.GetPubsResponse, error) {
	payload := make([]*models.GetPubsResponse, 0)
	data := new(models.GetPubsResponse)
	genreId, _ := strconv.ParseInt(id, 10, 64)
	row := p.db.QueryRow(`select * from publications where id = ? OR slug = ? `, genreId, id)
	err := row.Scan(
		&data.Id,
		&data.Name,
		&data.Meta.Slug,
		&data.Meta.FoundingDate,
		&data.Meta.Desc,
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
func (p *PubsRepo) Delete(ctx context.Context, id string) (bool, error) {
	pId, _ := strconv.ParseInt(id, 10, 64)
	res, err := p.db.Exec(`update publications set status = 0 where id=? OR slug = ?`, pId, id)
	rows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if rows >= 0 {
		return true, nil
	}
	return false, nil
}

func (p *PubsRepo) Update(ctx context.Context, r *models.UpdateReqPublish) (bool, error) {
	//insert record in database
	res, err := p.db.Exec(`Update publications set name = ?, status = ?, founding_date = ?, description = ?,  updated_at = ? where id = ?`, strings.TrimSpace(*r.Name), *r.Status, r.Meta.FoundingDate.Format(cons.MYSQL_DATE_FORMAT), *r.Meta.Desc, time.Now(), *r.Id)
	if err != nil {
		return false, err
	}
	row, _ := res.RowsAffected()
	if row >= 0 {
		return true, nil
	}
	return false, models.ErrNotFound
}

/**
 * [func description]
 * @param  {[type]} p [description]
 * @return {[type]}   [description]
 */

func (p *PubsRepo) generatePubsSlug(slug string) (string, error) {

	var id int32
	slug = strings.TrimSpace(slug)
	slug = strings.ToLower(slug)
	//change all spaces to - in order to create a valid Slug
	slug = strings.Replace(slug, " ", "-", -1)
	//check if slug already present in DB
	row := p.db.QueryRow(`SELECT id FROM publications WHERE slug=?`, slug)
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
