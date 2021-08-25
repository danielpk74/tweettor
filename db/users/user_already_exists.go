package users

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"github.com/danielpk74/tweettor/db"
	"github.com/danielpk74/tweettor/models"
)

// UserAlreadyExists Checks if an email already exists in the DB
func UserAlreadyExists(email string) (models.User, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"email": email}
	var result models.User

	err := db.Conn.TweettorCollection("users").FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false, ""
	}

	ID := result.ID.Hex()
	return result, true, ID
}
