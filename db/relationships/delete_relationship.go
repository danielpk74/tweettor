package relationships

import (
	"context"
	"github.com/danielpk74/tweettor/db"
	"github.com/danielpk74/tweettor/models"
	"time"
)

func DeleteRelationship(r models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := db.Conn.TweettorCollection("relationship")

	_, err := collection.DeleteOne(ctx, r)
	if err != nil {
		return false, err
	}

	return true, nil
}
