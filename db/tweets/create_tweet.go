package tweets

import (
	"context"
	"github.com/danielpk74/tweettor/db"
	"github.com/danielpk74/tweettor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CreateTweet(t models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	tweet := bson.M{
		"user_id": t.UserID,
		"tweet":   t.Tweet,
		"date":    t.Date,
	}

	result, err := db.Conn.TweettorCollection("tweets").InsertOne(ctx, tweet)
	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)

	return objId.String(), true, nil
}
