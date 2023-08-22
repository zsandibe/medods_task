package migration

import (
	"context"
	"fmt"
	"log"
	"task/internal/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *Db) Create(userToken *repository.UserToken) error {
	mongoRes, err := db.collection.InsertOne(context.Background(), userToken)
	if err != nil {
		log.Printf("error in creating: %v\n", err)
		return err
	}

	id, flag := mongoRes.InsertedID.(primitive.ObjectID)
	if !flag {
		log.Println(id.Hex())
		return err
	}
	log.Println(mongoRes.InsertedID)
	return nil
}

func (db *Db) Get(userGUID string, bindTokens string) (*repository.UserToken, error) {
	filter := bson.M{
		"user_guid":   userGUID,
		"bind_tokens": bindTokens,
	}
	fmt.Println(filter)
	mongoRes := db.collection.FindOne(context.Background(), filter)
	fmt.Println(mongoRes.Err())
	if mongoRes.Err() != nil {
		log.Printf("error in getting: %v\n", mongoRes.Err())
		return nil, mongoRes.Err()
	}

	userTokens := &repository.UserToken{}

	if err := mongoRes.Decode(&userTokens); err != nil {
		log.Printf("error in decoding: %v\n", err)
		return nil, err
	}
	return userTokens, nil
}

func (db *Db) Update(old *repository.UserToken, new *repository.UserToken) error {
	filter := bson.M{
		"refresh_token": bson.M{"$eq": old.RefreshToken},
		"bind_tokens":   bson.M{"$eg": old.BindTokens},
		"user_guid":     bson.M{"$eq": old.UserGUID},
	}

	update := bson.M{
		"$set": bson.M{
			"user_guid":     new.UserGUID,
			"bind_tokens":   new.BindTokens,
			"refresh_token": new.RefreshToken,
		},
	}
	mongoRes, err := db.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("error in updating: %v\n", err)
		return err
	}
	fmt.Println("updated", mongoRes)
	return nil
}
