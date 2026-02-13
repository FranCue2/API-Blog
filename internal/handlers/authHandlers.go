package handlers

import (
	"context"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tu-usuario/blog-api/internal/auth"
	db "github.com/tu-usuario/blog-api/internal/database"
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

	err = auth.RegisterUser(ctx, input.Email, input.Password)

	if err != nil {
		switch {
			case errors.Is(err, db.ErrUserAlreadyExists):
				c.JSON(409, gin.H{"error":"Ese email ya esta en uso"})

			case errors.Is(err, auth.ErrFailedToEncryptPassword):
				c.JSON(500, gin.H{"error":"Error interno de seguridad"})

			default:
				c.JSON(500, gin.H{"error":"Error interno del servidor"})
		}
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