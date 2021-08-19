package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Conn MongoDB Connection instance.
var Conn = newMongoClient()
var clientOptions = options.Client().ApplyURI("mongodb+srv://book_user:book_user@cluster0.8yb8h.mongodb.net/books_store?retryWrites=true&w=majority")

type MongoClient struct {
	Client *mongo.Client
}

// NewMongoClient connect our app to mongoDB
func newMongoClient() *MongoClient {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error Connecting To MongoDB: %s", err)
		return nil
	}

	//err = client.Ping(context.TODO(), nil)
	//if err != nil {
	//	log.Fatalf("Error Ping: %s", err)
	//	return nil
	//}

	log.Println("Connected to Mongo!")
	return &MongoClient{
		client,
	}
}

// CheckConnection check if our app has connection with mongoDB
func (c *MongoClient) CheckConnection() bool {
	err := c.Client.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}

	return true
}

// TweetorCollection check if our app has connection with mongoDB
func (c *MongoClient) TweetorCollection(collection string) *mongo.Collection {
	return c.Client.Database("tweetor").Collection(collection)
}
