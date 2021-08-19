package users

import (
	"context"
	"github.com/danielpk74/tweettor/db"
	"github.com/danielpk74/tweettor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type dbClient struct {
	client *db.MongoClient
}

// CreateUser Create a new user in the Database
func CreateUser(u *models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	u.Password, _ = EncryptPassword(u.Password)
	result, err := db.Conn.TweetorCollection("users").InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
