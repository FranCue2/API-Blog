package auth

import "errors"

var (
	ErrInvalidInput			   = errors.New("input was invalid")
	

	ErrFailedToEncryptPassword = errors.New("failled to ecrypt password")
	ErrFailedToProcessPassword = errors.New("failled to process password")

	ErrFailedToProcessToken    = errors.New("failled to process token")
	ErrFailedToGenerateToken   = errors.New("failled to generate token")
	ErrInvalidToken			   = errors.New("invalid token")
	ErrInvalidRole			   = errors.New("invalid role")
)
