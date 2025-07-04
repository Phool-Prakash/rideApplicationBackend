package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() *mongo.Client{
	client,err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("Mongo client error:",err)
	}

	ctx,cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err !=nil {
		log.Fatal("Mongo connection failed:",err)
	}

	log.Println("Database connected")
	DB = client
	return DB
}