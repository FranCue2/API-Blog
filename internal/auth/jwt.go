package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tu-usuario/blog-api/internal/models"
)


var jwtKey = []byte(os.Getenv("JWT_SECRET"))


type UserClaims struct{
	Role 	models.Role 	`json:"role"`
	Email	string			`json:"email"`

	jwt.RegisteredClaims
}

func GenerateToken(userId string, email string, role models.Role) (string ,error){
	if !role.IsValid(){
		return "", errors.New("rol invalido")
	}

	experationDate := time.Now().Add(24 * time.Hour)

	claims := &UserClaims{
		Role:	role,
		Email:	email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: userId,
			ExpiresAt: jwt.NewNumericDate(experationDate),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "blog-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*UserClaims, error){

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token)(interface{},error){return jwtKey, nil})

	if err != nil{
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token invalido")
}