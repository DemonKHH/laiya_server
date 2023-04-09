package serviceUser

import (
	"context"
	modelUser "laiya_server/internal/model/user"
	db "laiya_server/service/mongo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = db.OpenCollection(db.GetMongoClient(), "users")

func GetUser(userId string) (modelUser.User, error) {
	var user modelUser.User
	var err error
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = userCollection.FindOne(ctx, bson.M{"userid": userId}).Decode(&user)
	return user, err
}

func GetUsers() ([]bson.M, error) {
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var users []bson.M
	if err = cursor.All(ctx, &users); err != nil {
		log.Fatal(err)
	}
	return users, err
}
