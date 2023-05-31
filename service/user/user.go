package serviceUser

import (
	"context"
	"laiya_server/pkg/common/response"
	db "laiya_server/service/mongo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func UpdatePermissions(userId string, permissions []string) (response.LoginResponseWithUser, error) {
	var err error
	var user response.LoginResponseWithUser
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"userid": userId}
	// 构造更新操作
	update := bson.M{"$set": bson.M{"permissions": permissions}}
	// 执行更新操作，并返回更新后的文档
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = userCollection.FindOneAndUpdate(ctx, filter, update, options).Decode(&user)
	return user, err
}
