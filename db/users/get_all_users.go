package users

import (
	"context"
	"fmt"
	"github.com/danielpk74/tweettor/db"
	"github.com/danielpk74/tweettor/db/relationships"
	"github.com/danielpk74/tweettor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func GetAllUsers(ID string, page int64, searchString string, searchType string) ([]*models.AllUsersResponse, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	userCollection := db.Conn.TweettorCollection("users")

	var findOptions = options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	condition := bson.M{
		"name": bson.M{"$regex": `(?i)` + searchString},
	}

	users, err := userCollection.Find(ctx, condition, findOptions)
	if err != nil {
		_ = fmt.Errorf("there was an error %s", err.Error())
		return []*models.AllUsersResponse{}, false
	}

	var usersHaveRelation, include bool
	var results []*models.AllUsersResponse
	for users.Next(context.TODO()) {
		var u models.AllUsersResponse
		err := users.Decode(&u)
		if err != nil {
			_ = fmt.Errorf("there was an error %s", err.Error())
			return results, false
		}

		var r models.Relationship
		r.UserId = ID
		r.FollowedUserId = u.ID.Hex()

		include = false
		usersHaveRelation, err = relationships.FindRelationship(r)
		if searchType == "new" && usersHaveRelation == false {
			include = true
		}

		if searchType == "follow" && usersHaveRelation == true {
			include = true
		}

		if include {
			results = append(results, &u)
		}
	}

	return results, true
}
