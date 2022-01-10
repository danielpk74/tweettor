package tweets

import (
	"context"
	"github.com/danielpk74/tweettor/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func DeleteTweet(tweetID string, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	tweetCollection := db.Conn.TweettorCollection("tweets")

	tweetObjID, _ := primitive.ObjectIDFromHex(tweetID)
	condition := bson.M{"_id": tweetObjID, "user_id": userID}

	_, err := tweetCollection.DeleteOne(ctx, condition)
	return err
}
