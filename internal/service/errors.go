package service

import "errors"

var (
	errInvalidCredentials error = errors.New("invalid credentials")
	errUserNotFound error = errors.New("user not found")
)
