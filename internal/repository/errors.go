package repository

import "errors"

var (
	errUsernameIsAlreadyTaken error = errors.New("this username is already taken")
	errTodoNotFound error = errors.New("todo not found")
)