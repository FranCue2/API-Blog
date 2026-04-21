package auth

import (
	"context"
	"errors"
	"time"

	db "github.com/tu-usuario/blog-api/internal/database"
	"github.com/tu-usuario/blog-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin(email string, password string) error {

	if email == "" || password == "" {
		return ErrInvalidInput
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.FindUserCredentialsByEmail(ctx, email)

	if err == nil {
		return nil
	}

	if err != db.ErrUserDoesNotExist {
		return err
	}

	err = CreateAdmin(ctx, email, password)

	return err
}

func Login(ctx context.Context, email string, password string) (string, error) {

	userCred, err := db.FindUserCredentialsByEmail(ctx, email)

	if err != nil {
		return "", ErrInvalidInput
	}

	err = bcrypt.CompareHashAndPassword([]byte(userCred.PasswordHash), []byte(password))

	if err != nil {
		if(errors.Is(err, bcrypt.ErrMismatchedHashAndPassword)) {return "", ErrInvalidInput}
		return "", ErrFailedToProcessPassword
	}

	return GenerateToken(userCred.ID.Hex(), userCred.Email, userCred.Role)

}

func CreateAdmin(ctx context.Context, email string, password string) error {
	return generateAndInsertCredentials(ctx, email, password, models.RoleAdmin)
}

func RegisterUser(ctx context.Context, email string, password string) error {
	return generateAndInsertCredentials(ctx, email, password, models.RoleUser)
}

func generateAndInsertCredentials(ctx context.Context, email string, password string, role models.Role) error {

	passwordHash, err := EncryptPassword(password)

	if err != nil {
		return ErrFailedToEncryptPassword
	}

	creds := models.UserCredentials{
		Role:         role,
		Email:        email,
		PasswordHash: passwordHash,
		TOTPEnabled:  false,
	}

	_, err = db.InsertUserCredentials(ctx, &creds)

	if err != nil {
		return err
	}

	return nil
}

func EncryptPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(passwordHash), err
}
