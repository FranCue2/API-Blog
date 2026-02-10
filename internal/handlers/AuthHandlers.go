package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tu-usuario/blog-api/internal/auth"
	"github.com/tu-usuario/blog-api/internal/constants"
	db "github.com/tu-usuario/blog-api/internal/database"
	"github.com/tu-usuario/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type loginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type registerInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context) {

	var input registerInput

	err := c.ShouldBindBodyWithJSON(&input)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not good enough information", "message": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	passwordHash, err := auth.EncryptPassword(input.Password)

	if err != nil {
		error := "Fallo al generar contrasena"
		c.JSON(500, gin.H{"error": error})
		return
	}

	creds := models.UserCredentials{
		Role:         models.RoleUser,
		Email:        input.Email,
		PasswordHash: passwordHash,
		TOTPEnabled:  false,
	}

	_, err = db.GetCollection(constants.AuthCredentialsCollections).InsertOne(ctx, creds)

	if err != nil {
		var error = fmt.Sprintf("error desconocido al guardar %v", err)
		if mongo.IsDuplicateKeyError(err) {
			error = "este email ya esta en uso"
		}

		c.JSON(500, gin.H{"error": error})
		return
	}

	c.JSON(200, "Creado usuario con exito")
}

func Login(c *gin.Context) {
	var input loginInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Datos de entrada invalidos: Email y Contrasena obligatorios"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	token, err := auth.Login(ctx, input.Email, input.Password)

	if err != nil {
		c.JSON(401, gin.H{"error": "Email y/o Conreasena equivocada", "mensage": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})

}

func CreateAdmin(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	password := "admin"
	passwordHash, err := auth.EncryptPassword(password)

	if err != nil {
		error := "Fallo al generar contrasena"
		c.JSON(500, gin.H{"error": error})
		return
	}

	creds := models.UserCredentials{
		Email:        "admin@example.com",
		Role:         models.RoleAdmin,
		PasswordHash: passwordHash,
		TOTPEnabled:  false,
	}

	_, err = db.GetCollection(constants.AuthCredentialsCollections).InsertOne(ctx, creds)

	if err != nil {
		var error = fmt.Sprintf("error desconocido al guardar %v", err)
		if mongo.IsDuplicateKeyError(err) {
			error = "ya esta el admin"
		}

		c.JSON(500, gin.H{"error": error})
		return
	}

	c.JSON(200, "Creado admin con exito")
}
