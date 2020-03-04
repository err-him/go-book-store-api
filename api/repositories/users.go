package repositories

import (
	c "book-store-api/api/constants"
	"book-store-api/api/models"
	"book-store-api/api/utils"
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"strings"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) models.UsersRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) Create(ctx context.Context, r *models.User) (*models.User, error) {

	//generate api Key for user
	uKey := utils.GenerateRandomStr(30)

	//generate password salt
	uPass, err := generateUserPass(*r.Password)
	if err != nil {
		return nil, err
	}
	//insert record in database
	res, err := u.db.Exec(`Insert into user (name, password, api_key) value (?, ?, ?)`, strings.TrimSpace(strings.ToLower(*r.Name)), uPass, uKey)

	if err != nil {
		return nil, err
	}
	insertedId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	payload := &models.User{
		Id:       insertedId,
		UserName: r.Name,
	}
	return payload, nil
}

func (u *UserRepo) Verify(ctx context.Context, r *models.User) (*models.User, error) {
	//generate api Key for user
	uPass, err := generateUserPass(*r.Password)
	if err != nil {
		return nil, err
	}
	data := &models.User{}
	row := u.db.QueryRow(`select id,name, api_key from user where name = ? AND password = ? `, strings.TrimSpace(strings.ToLower(*r.Name)), uPass)
	err = row.Scan(
		&data.Id,
		&data.Name,
		&data.Token,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrInvalidCredential
		}
		return nil, err
	}
	//get user JWT Token
	data.Token, err = utils.IssueJWTToken(data.Id, *data.Name, data.Token)
	return data, err
}

func generateUserPass(pass string) (string, error) {
	//get user password salt

	pSalt, err := utils.GetEnvVar(c.PASSWORD_SALT_KEY)
	if err != nil {
		return "", err
	}
	newPass := pass + pSalt
	//generate hash for new password
	hash := md5.Sum([]byte(newPass))
	return hex.EncodeToString(hash[:]), nil
}
