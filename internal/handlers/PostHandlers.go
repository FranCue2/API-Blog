package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tu-usuario/blog-api/internal/database"
	"github.com/tu-usuario/blog-api/internal/format"
	"go.mongodb.org/mongo-driver/v2/bson"
)


func CreatePost(c *gin.Context) {
	
	var post format.PostFormat
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	ctx , cancel:= context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	
	res, err := db.Client.Database("Blog_DB").Collection("posts").InsertOne(ctx, post)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al crear la publicación"})
		return
	}

	id, ok := res.InsertedID.(bson.ObjectID)
	if !ok {
		c.JSON(500, gin.H{"error": "Error al obtener el ID de la publicación"})
		return
	}

	c.JSON(201, gin.H{"message": "Publicación"+ post.Title + "creada exitosamente, con id:" + id.Hex()})
}

func RetreiveAllPosts(c *gin.Context) {
	ctx , cancel:= context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	res, err := db.Client.Database("Blog_DB").Collection("posts").Find(ctx, bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener las publicaciones"})
		return
	}

	posts := []format.PostFormat{}

	if err := res.All(c.Request.Context(), &posts); err != nil {
		c.JSON(500, gin.H{"error": "Error al procesar las publicaciones"})
		return
	}


	c.JSON(200, gin.H{"posts": posts})
}