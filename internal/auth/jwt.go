package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tu-usuario/blog-api/internal/constants"
	"github.com/tu-usuario/blog-api/internal/models"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type UserClaims struct {
	Role  models.Role `json:"role"`
	Email string      `json:"email"`

	jwt.RegisteredClaims
}

func GenerateToken(userId string, email string, role models.Role) (string, error) {
	if !role.IsValid() {
		return "", ErrInvalidRole
	}

	experationDate := time.Now().Add(constants.TokenExperitationTime)

	claims := &UserClaims{
		Role:  role,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId,
			ExpiresAt: jwt.NewNumericDate(experationDate),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "blog-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", ErrFailedToGenerateToken
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*UserClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) { return jwtKey, nil })
	fmt.Printf("err: %v\n", err)
	if err != nil {
		return nil, ErrFailedToProcessToken
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
