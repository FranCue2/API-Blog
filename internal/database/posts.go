package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/tu-usuario/blog-api/internal/constants"
	"github.com/tu-usuario/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func InsertPost(ctx context.Context, post models.PostModel) (bson.ObjectID, error) {
	id, err := insertOneIntoCollection(ctx, post, constants.PostsCollections)
	if err != nil {
		switch{
		case errors.Is(err, ErrFailedToInsert):
			return bson.NilObjectID, ErrFailedToInsertPost
		default:
			return bson.NilObjectID, err
		}
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
	filter := bson.M{}

	if titleQuery != ""{
		filter["title"] = bson.M{
			"$regex":   titleQuery,
			"$options": "i",
		}
	}

	if authorQuery != ""{
		filter["author"] = bson.M{
			"$regex":   authorQuery,
			"$options": "i",
		}
	}

	if published_atQuery != ""{
		filter["published_at"] = bson.M{
			"$regex":   published_atQuery,
			"$options": "i",
		}
	}

	return findPostsWithFilter(ctx, filter)
}

func FindAllPosts(ctx context.Context) (*[]models.PostModel, error) {

	filter := bson.M{}

	return findPostsWithFilter(ctx, filter)

}

func findOnePostWithFilter(ctx context.Context, filter bson.M) (*models.PostModel, error) {
	fmt.Printf("filter: %v\n", filter)
	
	posts := models.PostModel{}
	err := findOneWithFilterFromColletionOfType(ctx, filter, constants.PostsCollections, &posts)

	if err != nil {
		switch{
		case errors.Is(err, ErrFailedToFind):
			return nil, ErrFailedToFindPosts
		case errors.Is(err, ErrFailedToProccess):
			return nil, ErrFailedToProccessPosts
		default:
			return nil, err
		}
	}
	return &posts, nil

}

func findPostsWithFilter(ctx context.Context, filter bson.M) (*[]models.PostModel, error) {

	posts := []models.PostModel{}
	err := findWithFilterFromCollectionOfType(ctx, filter, constants.PostsCollections, &posts)

	if err != nil {
		switch{
		case errors.Is(err, ErrFailedToFind):
			return nil, ErrFailedToFindPosts
		case errors.Is(err, ErrFailedToProccess):
			return nil, ErrFailedToProccessPosts
		default:
			return nil, err
		}
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

	err := deleteOneWithFilterFromCollection(ctx, filter, constants.PostsCollections)

	if err != nil {
		return ErrFailedToDeleteOnePost
	}

	return nil
}

