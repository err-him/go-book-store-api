package models

import (
	"context"
	"time"
)

type User struct {
	Id        int64      `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	UserName  *string    `json:"username,omitempty"`
	Password  *string    `json:"password,omitempty"`
	Status    *bool      `json:"status,omitempty"`
	Token     string     `json:"token,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type UsersRepo interface {
	Create(ctx context.Context, r *User) (*User, error)
	Verify(ctx context.Context, r *User) (*User, error)
}
