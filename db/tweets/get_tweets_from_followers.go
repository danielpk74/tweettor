package tweets

import (
	"context"
	"github.com/danielpk74/tweettor/db"
	"github.com/danielpk74/tweettor/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func GetTweetsFromFollowers(userID string, page int) ([]models.TweetFromFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	tweetCollection := db.Conn.TweettorCollection("relationship")
	skip := (page - 1) * 20
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"user_id": userID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "followed_user_id",
			"foreignField": "user_id",
			"as":           "tweet",
		},
	})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := tweetCollection.Aggregate(ctx, conditions)
	var results []models.TweetFromFollowers
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return results, false
	}

	return results, true
}
