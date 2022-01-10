package users

import (
	"context"
	"github.com/danielpk74/tweettor/db"
	"github.com/danielpk74/tweettor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func ModifyProfile(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	userCollection := db.Conn.TweettorCollection("users")

	profile := make(map[string]interface{})
	if len(u.Name) > 0 {
		profile["name"] = u.Name
	}

	if len(u.LastName) > 0 {
		profile["last_name"] = u.LastName
	}

	if !u.BirthDay.IsZero() {
		profile["birthDay"] = u.BirthDay
	}

	if len(u.Avatar) > 0 {
		profile["avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		profile["banner"] = u.Banner
	}

	if len(u.Biography) > 0 {
		profile["Biography"] = u.Biography
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	newProfileData := bson.M{"$set": profile}
	condition := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := userCollection.UpdateOne(ctx, condition, newProfileData)
	if err != nil {
		return false, err
	}

	return true, nil
}
