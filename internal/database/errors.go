package db

import "errors"

var (
	ErrFailedToConnectToDataBase = errors.New("failed to connect to database")

	ErrFailledToCreateIndexis 	 = errors.New("failed to create indexis succsesfuly")

	ErrUserAlreadyExists 		 = errors.New("user already exists in database")
	ErrUserDoesNotExist			 = errors.New("user does not exist yet")
)
