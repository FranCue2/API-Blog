package db

import "errors"

var (
	ErrUserAlreadyExists 	= errors.New("User Already Exists In Database")
	ErrUserDoesNotExist		= errors.New("User does not exist yet")
)
