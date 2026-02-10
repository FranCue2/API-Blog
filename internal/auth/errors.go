package auth

import "errors"

var (
	ErrInvalidInput			   = errors.New("input was invalid")
	
	ErrFailedToEncryptPassword = errors.New("failled to ecrypt password")

)
