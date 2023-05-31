package serviceUser

import (
	"context"
	"laiya_server/pkg/common/response"
	db "laiya_server/service/mongo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = db.OpenCollection(db.GetMongoClient(), "users")

func GetUser(userId string) (response.LoginResponse, error) {
	var user response.LoginResponse
	var err error
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = userCollection.FindOne(ctx, bson.M{"userid": userId}).Decode(&user)
	return user, err
}

func GetUsers() ([]response.LoginResponseWithUser, error) {
	var users []response.LoginResponseWithUser
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &users); err != nil {
		log.Fatal(err)
	}
	return users, err
}
