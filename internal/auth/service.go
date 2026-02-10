package auth

import (
	"context"
	"errors"

	db "github.com/tu-usuario/blog-api/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx context.Context, email string, password string) (string, error){	

	userCred, err := db.FindUserCredentialsByEmail(ctx,email)

	if err!=nil{
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userCred.PasswordHash), []byte(password))

	if err!= nil{
		return "", errors.New("contrasena incorrecta")
	}

	return GenerateToken(userCred.ID.Hex(), userCred.Email, userCred.Role)

}

func EncryptPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(passwordHash), err
}