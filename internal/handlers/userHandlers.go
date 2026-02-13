package handlers

import (
	"context"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tu-usuario/blog-api/internal/database"
)

func GetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	usersCreds, err := db.FindAllUsers(ctx, c)

	if err != nil{
		switch{
		case errors.Is(err, db.ErrFailedToFindUser):
			c.JSON(500, gin.H{"error": "Se fallo al buscar a los usuarios"})
		case errors.Is(err, db.ErrFailedToProccessUser):
			c.JSON(500, gin.H{"error":"Se fallo al procesar a los usuarios"})
		default:
			c.JSON(500, gin.H{"error": "Error deconocido: " + err.Error()})
		}

		return
	}

	c.JSON(200, gin.H{"credentials": *usersCreds})
}
