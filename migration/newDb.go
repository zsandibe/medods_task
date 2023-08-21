package migration

import (
	"context"
	"log"
	"task/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDatabase(config config.Config) *mongo.Database {
	port := config.MongoURL

	client, err := mongo.NewClient(options.Client().ApplyURI(port))
	if err != nil {
		log.Printf("Error connecting to MongoDB: %v", err)
		return nil
	}

	if err = client.Connect(context.TODO()); err != nil {
		log.Printf("client MongoDB: %v\n", err)
		return nil
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Printf("client connecting to MongoDB: %v\n", err)
		return nil
	}
	log.Println("Connected to MongoDB")
	return client.Database(config.NameDB)
}
