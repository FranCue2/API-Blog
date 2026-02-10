package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tu-usuario/blog-api/internal/constants"
	db "github.com/tu-usuario/blog-api/internal/database"
	"github.com/tu-usuario/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	res, err := db.GetCollection(constants.AuthCredentials).Find(ctx, bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener las publicaciones"})
		return
	}

	usersCreds := []models.UserCredentials{}

	if err := res.All(c.Request.Context(), &usersCreds); err != nil {
		c.JSON(500, gin.H{"error": "Error al procesar las publicaciones"})
		return
	}

	c.JSON(200, gin.H{"credentials": usersCreds})
}

