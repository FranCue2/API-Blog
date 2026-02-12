package db

import (
	"context"

	"github.com/tu-usuario/blog-api/internal/constants"
	"github.com/tu-usuario/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func InsertPost(ctx context.Context, post models.PostModel) (bson.ObjectID, error) {
	res, err := GetCollection(constants.PostsCollections).InsertOne(ctx, post)

	if err != nil {
		return bson.NilObjectID, ErrFailedToInsertPost
	}

	id, ok := res.InsertedID.(bson.ObjectID)

	if !ok {
		return bson.NilObjectID, ErrFailedToObtainID
	}

	return id, nil
}

func FindPostWithID(ctx context.Context, id string) (*models.PostModel, error) {
	idObj, err := getObjectId(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": idObj}
	return findOnePostWithFilter(ctx, filter)
}

func FindPostsWithQuery(ctx context.Context, titleQuery string, authorQuery string, published_atQuery string) (*[]models.PostModel, error) {
	filter := bson.M{
		"title": bson.M{
			"$regex":   titleQuery,
			"$options": "i",
		},
		"author": bson.M{
			"$regex":   authorQuery,
			"$options": "i",
		},
		"published_at": bson.M{
			"$regex":   published_atQuery,
			"$options": "i",
		},
	}

	return findPostsWithFilter(ctx, filter)
}

func FindAllPosts(ctx context.Context) (*[]models.PostModel, error) {

	filter := bson.M{}

	return findPostsWithFilter(ctx, filter)

}

func findOnePostWithFilter(ctx context.Context, filter bson.M) (*models.PostModel, error) {
	res := GetCollection(constants.PostsCollections).FindOne(ctx, filter)

	if err := res.Err(); err != nil {
		return nil, ErrFailedToFindPosts
	}

	var post models.PostModel

	if err := res.Decode(&post); err != nil {
		return nil, ErrFailedToProccessPosts
	}

	return &post, nil

}

func findPostsWithFilter(ctx context.Context, filter bson.M) (*[]models.PostModel, error) {
	res, err := GetCollection(constants.PostsCollections).Find(ctx, filter)

	if err != nil {
		return nil, ErrFailedToFindPosts
	}

	posts := []models.PostModel{}

	if err := res.All(ctx, &posts); err != nil {
		return nil, ErrFailedToProccessPosts
	}

	return &posts, nil
}

func DeleteWithID(ctx context.Context, id string) error {
	idObj, _ := getObjectId(id)

	filter := bson.M{"_id": idObj}

	return deleteOnePostWithFilter(ctx, filter)

}

func EmptyPosts() error {
	return emptyCollection(constants.PostsCollections)
}

func deleteOnePostWithFilter(ctx context.Context, filter bson.M) error {

	_, err := GetCollection(constants.PostsCollections).DeleteOne(ctx, filter)

	if err != nil {
		return ErrFailedToDeleteOnePost
	}

	return nil
}
