package db

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	ConnectDB()
}

func ConnectDB() {
	var err error
	s := flag.String("env", "dev", "环境变量")
	flag.Parse()
	if *s == "dev" {
		err = godotenv.Load(".env.dev")
		if err != nil {
			log.Fatal("error loading .env.dev file")
		}
	} else {
		err = godotenv.Load(".env.prod")
		if err != nil {
			log.Fatal("error loading .env.prod file")
		}
	}

	MongoDb := os.Getenv("MONGODB_URL")
	log.Printf("mongo url: %s", MongoDb)
	client, err = mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Print("Connection failed to MongoDB")
		log.Fatal(err)
	}

	fmt.Print("Connected to MongoDB")
}

func GetMongoClient() *mongo.Client {
	return client
}

func OpenCollection(client *mongo.Client, CollectionName string) *mongo.Collection {
	database := os.Getenv("MONGODB_DATABASE")
	return client.Database(database).Collection(CollectionName)
}
