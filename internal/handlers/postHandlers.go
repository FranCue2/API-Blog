package handlers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tu-usuario/blog-api/internal/database"
	"github.com/tu-usuario/blog-api/internal/models"
)

func CreatePost(c *gin.Context) {

	var post models.PostModel
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	post.Author = getUserName(c) //Cargamos el nombre del usuario que creo el post
	post.PublishedAt = time.Now()

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	id, err := db.InsertPost(ctx, post)
	if err != nil {
		switch {
		case errors.Is(err, db.ErrFailedToInsertPost):
			c.JSON(500, gin.H{"error": "Error al crear la publicación"})

		case errors.Is(err, db.ErrFailedToObtainID):
			c.JSON(500, gin.H{"error": "Error al obtener el ID de la publicación"})

		default:
			c.JSON(500, gin.H{"error": "Error desconocido: " + err.Error()})
		}
		return
	}

	c.JSON(201, gin.H{"message": "Publicación " + post.Title + " creada exitosamente", "id:": id.Hex()})
}

func RetreiveAllPosts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	posts, err := db.FindAllPosts(ctx)

	if err != nil {

		switch {
		case errors.Is(err, db.ErrFailedToFindPosts):
			c.JSON(500, gin.H{"error": "Error al obtener las publicaciones"})
		case errors.Is(err, db.ErrFailedToProccessPosts):
			c.JSON(500, gin.H{"error": "Error al procesar las publicaciones"})
		default:
			c.JSON(500, gin.H{"error": "Error desconocido: " + err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"posts": posts})
}

func GetPostByID(c *gin.Context) {

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	post, err := db.FindPostWithID(ctx, id)

	if err != nil {
		switch {
		case errors.Is(err, db.ErrFailedToFindPosts):
			c.JSON(500, gin.H{"error": "Error al obtener las publicaciones"})

		case errors.Is(err, db.ErrFailedToProccessPosts):
			c.JSON(500, gin.H{"error": "Error al procesar las publicaciones"})

		case errors.Is(err, db.ErrFailedToProcesID):
			c.JSON(500, gin.H{"error": "Error al procesar el ID"})

		default:
			c.JSON(500, gin.H{"error": "Error desconocido: " + err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"posts": post})
}

func SearchPosts(c *gin.Context) {
	titleQuery := c.Query("title")
	authorQuery := c.Query("author")
	published_atQuery := c.Query("published_at")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	posts, err := db.FindPostsWithQuery(ctx, titleQuery, authorQuery, published_atQuery)

	if err != nil {

		switch {
		case errors.Is(err, db.ErrFailedToFindPosts):
			c.JSON(500, gin.H{"error": "Error al obtener las publicaciones"})
		case errors.Is(err, db.ErrFailedToProccessPosts):
			c.JSON(500, gin.H{"error": "Error al procesar las publicaciones"})
		default:
			c.JSON(500, gin.H{"error": "Error desconocido: " + err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"posts": posts})
}

func DeleteByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	id, found := c.GetQuery("id")

	if !found {
		c.JSON(500, gin.H{"error": "ID missing"})
	}

	err := db.DeleteWithID(ctx, id)

	if err != nil {
		error := fmt.Sprintf("Error al borrar post de id: %s", id)
		c.JSON(500, gin.H{"error": error})
	}

	mensaje := fmt.Sprintf("Eliminado con exito al post de id: %s", id)
	c.JSON(200, gin.H{"message": mensaje})
}

func DeleteAllPosts(c *gin.Context) {
	err := db.EmptyPosts()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al eliminar todas las publicaciones"})
		return
	}
	c.JSON(200, gin.H{"message": "Todas las publicaciones han sido eliminadas exitosamente"})
}
