package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostModel struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string        `json:"title" bson:"title"`
	Content     string        `json:"content" bson:"content"`
	Author      string        `json:"author" bson:"author"`
	PublishedAt string        `json:"published_at" bson:"published_at"`
}

