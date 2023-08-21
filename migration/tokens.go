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
		"userGUID":   userGUID,
		"bindTokens": bindTokens,
	}

	mongoRes := db.collection.FindOne(context.Background(), filter)

	if mongoRes.Err() != nil {
		log.Printf("error in getting: %v\n", mongoRes.Err())
		return nil, mongoRes.Err()
	}

	var userTokens repository.UserToken

	if err := mongoRes.Decode(&userTokens); err != nil {
		log.Printf("error in decoding: %v\n", err)
		return nil, err
	}
	return &userTokens, nil
}

func (db *Db) Update(old *repository.UserToken, new *repository.UserToken) error {
	filter := bson.M{
		"refreshToken": bson.M{"$eq": old.RefreshToken},
		"bindTokens":   bson.M{"$eg": old.BindTokens},
		"userGUID":     bson.M{"$eq": old.UserGUID},
	}

	update := bson.M{
		"$set": bson.M{
			"userGUID":     new.UserGUID,
			"bindTokens":   new.BindTokens,
			"refreshToken": new.RefreshToken,
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
