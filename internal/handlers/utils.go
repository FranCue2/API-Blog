package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func getObjectId(c *gin.Context) (bson.ObjectID) {
	id := c.Param("id")

	idObj, err := bson.ObjectIDFromHex(id)

	if err != nil {
		error := fmt.Sprintf("fallo al procesar id: %s", err)
		c.JSON(500, gin.H{"error": error})
	}
	return idObj
}