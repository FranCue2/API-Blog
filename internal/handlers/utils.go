package handlers

import (
	"github.com/gin-gonic/gin"
)

//TODO falta hacer la logica de perfil para tomar el nombre de forma correcta
func getUserName(c *gin.Context) string{ 
	return c.GetString("userID")
}