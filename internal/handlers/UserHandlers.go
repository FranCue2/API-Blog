package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tu-usuario/blog-api/internal/database"
	"github.com/tu-usuario/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {

	email := c.Query("email")
	password := c.Query("password")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	passwordHash, err := encryptPassword(password)

	if err != nil {
		error := "Fallo al generar contrasena"
		c.JSON(500, gin.H{"error": error})
		return
	}

	creds := models.User_Credentials{
		Email:        email,
		PasswordHash: passwordHash,
		TOTPEnabled:  false,
	}

	_, err = db.GetCollection("auth_credentials").InsertOne(ctx, creds)

	if err != nil {
		var error = fmt.Sprintf("error desconocido al guardar %v", err)
		if mongo.IsDuplicateKeyError(err){
			error = "este email ya esta en uso"
		}

		c.JSON(500, gin.H{"error" : error})
		return
	}

	c.JSON(200, "Creado usuario con exito")
}

func GetUsers(c* gin.Context){
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	res, err := db.GetCollection("auth_credentials").Find(ctx, bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener las publicaciones"})
		return
	}

	posts := []models.User_Credentials{}

	if err := res.All(c.Request.Context(), &posts); err != nil {
		c.JSON(500, gin.H{"error": "Error al procesar las publicaciones"})
		return
	}

	c.JSON(200, gin.H{"posts": posts})
}

func encryptPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(passwordHash), err
}
