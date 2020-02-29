package models

import "book-store-api/api/utils"

type AuthorCreateReq struct {
	Id    *int64       `json:"id,omitempty"`
	About *AuthorAbout `json:"about"`
}

type AuthorAbout struct {
	YearsActive *utils.MysqlFormatDate `json:"years_active,omitempty"`
	Language    *string                `json:"language"`
	Personal    *AuhtorPersonal        `json:"personal"`
}

type AuhtorPersonal struct {
	Dob  *utils.MysqlFormatDate `json:"dob,omitempty"`
	Info *string                `json:"info"`
}

type AuthorRepo interface {
}
