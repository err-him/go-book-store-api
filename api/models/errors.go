package models

import "errors"

var (
	ErrNotFound          = errors.New("Requested item is not found!")
	ErrAlreadyPresent    = errors.New("Data already present")
	ErrInvalidCredential = errors.New("Unauthorized, Invalid username or password")
)
