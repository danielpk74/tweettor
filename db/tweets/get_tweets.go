package tweets

import (
	"context"
	"github.com/danielpk74/tweettor/db"
	"github.com/danielpk74/tweettor/models"
	"go.mongodb.org/mongo-driver/bson"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetTweetsPerUser(userID string, page int64, tweetsPerPage int64) ([]*models.Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	tweetsCollection := db.Conn.TweettorCollection("tweets")

	var userTweets []*models.Tweet
	condition := bson.M{"user_id": userID}

	options := options2.Find()
	options.SetLimit(tweetsPerPage)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * tweetsPerPage)

	tweets, err := tweetsCollection.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return userTweets, false
	}

	for tweets.Next(context.TODO()) {
		var tweet models.Tweet
		err := tweets.Decode(&tweet)
		if err != nil {
			log.Fatal(err.Error())
			return userTweets, false
		}

		userTweets = append(userTweets, &tweet)
	}

	return userTweets, true
}
