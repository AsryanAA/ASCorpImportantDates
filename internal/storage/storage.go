package storage

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrUsersIsNotExist = errors.New("'users is not exists")
)
