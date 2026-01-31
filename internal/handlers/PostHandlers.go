package handlers

import (
	"context"
	"fmt"
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

func GetPostByID(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	idObj := getObjectId(c)

	res := db.Client.Database("Blog_DB").Collection("posts").FindOne(ctx, bson.M{"_id": idObj})
	if err := res.Err(); err != nil {
		//c.JSON(500, gin.H{"error": "Error al obtener las publicaciones"})
		c.JSON(500, gin.H{"error": res.Err().Error()})
		return
	}

	var post format.PostFormat
	var err = res.Decode(&post)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error al procesar las publicaciones"})
		return
	}


	c.JSON(200, gin.H{"posts": post})
}


func SearchPosts(c *gin.Context){
	titleQuery 		  := c.Query("title")
	authorQuery		  := c.Query("author")
	published_atQuery := c.Query("published_at")

	filter := bson.M{
        "title": bson.M{
            "$regex": titleQuery, 
            "$options": "i",
        },
		"author": bson.M{
            "$regex": authorQuery, 
            "$options": "i",
        },
		"published_at": bson.M{
			"$regex":published_atQuery,
			"$options": "i",
		},
    }

	ctx , cancel:= context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	res, err := db.Client.Database("Blog_DB").Collection("posts").Find(ctx, filter)
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

func DeleteByID(c *gin.Context){
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	idObj := getObjectId(c)

	res, err := db.Client.Database("Blog_DB").Collection("posts").DeleteOne(ctx, bson.M{"_id": idObj})

	if err != nil {
		error := fmt.Sprintf("Error al borrar post de id %s, con error \n %d", &idObj, &err)
		c.JSON(500, gin.H{"error": error})
	}

	if res.DeletedCount == 0 {
		mensaje := fmt.Sprintf("No existe post con id %s", &idObj)
		c.JSON(200, gin.H{"mensaje":mensaje})
	}

	mensaje := fmt.Sprintf("Eliminado con exito al post de id %s", &idObj)
	c.JSON(200, gin.H{"message": mensaje})
}

func DeleteAllPosts(c *gin.Context) {
	err := db.EmptyDB()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al eliminar todas las publicaciones"})
		return
	}
	c.JSON(200, gin.H{"message": "Todas las publicaciones han sido eliminadas exitosamente"})
}

func getObjectId(c *gin.Context) (bson.ObjectID) {
	id := c.Param("id")

	idObj, err := bson.ObjectIDFromHex(id)

	if err != nil {
		error := fmt.Sprintf("fallo al procesar id: %s", err)
		c.JSON(500, gin.H{"error": error})
	}
	return idObj
}