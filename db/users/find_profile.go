package users

import (
	"context"
	"fmt"
	"github.com/danielpk74/tweettor/db"
	"github.com/danielpk74/tweettor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func FindProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Second)
	defer cancel()

	usersCollection := db.Conn.TweettorCollection("users")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{"_id": objID}

	err := usersCollection.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Record not found" + err.Error())
		return profile, err
	}

	return profile, nil
}
