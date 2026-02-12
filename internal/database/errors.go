package db

import "errors"

var (
	ErrFailedToConnectToDataBase = errors.New("failed to connect to database")

	ErrFailedToCreateIndexis = errors.New("failed to create indexis succsesfuly")

	ErrFailedToFindPosts     = errors.New("failed to finde posts")
	ErrFailedToProccessPosts = errors.New("failed to proccess posts")

	ErrFailedToProcesID = errors.New("failed to proccess id")

	ErrUserAlreadyExists = errors.New("user already exists in database")
	ErrUserDoesNotExist  = errors.New("user does not exist yet")

	ErrFailedToInsertPost = errors.New("failed to insert post")
	ErrFailedToObtainID = errors.New("failed to obtain ID")

	ErrFailedToDeleteOnePost = errors.New("failed deleating the post")
)
