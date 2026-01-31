package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tu-usuario/blog-api/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/posts", handlers.CreatePost)

	r.GET("/posts", handlers.RetreiveAllPosts)

	r.GET("/posts/:id", handlers.GetPostByID)

	r.GET("/posts/search", handlers.SearchPosts)

	r.DELETE("/posts", handlers.DeleteAllPosts)

	return r
}