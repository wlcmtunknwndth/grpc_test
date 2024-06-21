package storage

import "errors"

var (
	ErrUserExits    = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound  = errors.New("app not found")
)
