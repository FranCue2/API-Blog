package handlers

import (
	"github.com/gin-gonic/gin"
	db "github.com/tu-usuario/blog-api/internal/database"
)

func DeleteAllPosts(c *gin.Context) {
	err := db.EmptyDB()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al eliminar todas las publicaciones"})
		return
	}
	c.JSON(200, gin.H{"message": "Todas las publicaciones han sido eliminadas exitosamente"})
}