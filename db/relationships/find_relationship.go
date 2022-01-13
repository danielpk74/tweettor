package relationships

import (
	"context"
	"github.com/danielpk74/tweettor/db"
	"github.com/danielpk74/tweettor/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func FindRelationship(r models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := db.Conn.TweettorCollection("relationship")
	condition := bson.M{
		"user_id":          r.UserId,
		"followed_user_id": r.FollowedUserId,
	}

	var result models.Relationship

	err := collection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}

	return true, nil
}
