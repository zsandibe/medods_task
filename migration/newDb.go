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
		log.Fatal()
	}

	if err = client.Connect(context.TODO()); err != nil {
		log.Printf("client MongoDB: %v\n", err)
		log.Fatal()
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Printf("client connection to MongoDB: %v\n", err)
		log.Fatal()
	}
	log.Println("Connected to MongoDB")
	return client.Database(config.NameDB)
}
