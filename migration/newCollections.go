package migration

import (
	"task/internal/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type Db struct {
	collection *mongo.Collection
}

type Storage interface {
	Get(userGUID string, bindTokens string) (*repository.UserToken, error)
	Update(oldToken *repository.UserToken, newToken *repository.UserToken) error
	Create(userToken *repository.UserToken) error
}

func NewCollections(mongo *mongo.Database, collection string) *Db {
	return &Db{
		collection: mongo.Collection(collection),
	}
}

//
