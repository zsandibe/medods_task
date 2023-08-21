package migration

import (
	"task/internal/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
}

type Storage interface {
	Create(userToken *repository.UserToken) error
	Get(userGUID string, bindTokens string) (*repository.UserToken, error)
	Update(oldToken *repository.UserToken, newToken *repository.UserToken) error
}

func NewCollections(mongo *mongo.Database, collection string) *db {
	return &db{
		collection: mongo.Collection(collection),
	}
}

//
