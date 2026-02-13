package db

import "errors"

var (
	ErrFailedToConnectToDataBase = errors.New("failed to connect to database")

	ErrFailedToCreateIndexis = errors.New("failed to create indexis succsesfuly")

	///////////////Fails Related to looking up things
	ErrFailedToFind     = errors.New("failed to find")
	ErrFailedToProccess = errors.New("failed to proccess")

	ErrFailedToFindPosts     = errors.New("failed to finde posts")
	ErrFailedToProccessPosts = errors.New("failed to proccess posts")

	ErrFailedToFindUser     = errors.New("failed to find users")
	ErrFailedToProccessUser = errors.New("failed to proccess users")

	ErrFailedToProcesID = errors.New("failed to proccess id")

	ErrUserAlreadyExists = errors.New("user already exists in database")
	ErrUserDoesNotExist  = errors.New("user does not exist yet")

	ErrDuplicatedKey  = errors.New("key already in use")
	ErrFailedToInsert = errors.New("failed to insert")
	ErrDoesNotExist   = errors.New("it does not exist")

	ErrFailedToInsertUser   = errors.New("failed to insert user")
	ErrDuplicatedKeyForUser = errors.New("key dor user already in use")

	ErrFailedToInsertPost = errors.New("failed to insert post")
	ErrFailedToObtainID   = errors.New("failed to obtain id")

	ErrFailedToDeleteOnePost = errors.New("failed deleating the post")
)
