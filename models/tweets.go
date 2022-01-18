package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tweet struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID string             `bson:"user_id" json:"user_id,omitempty"`
	Tweet  string             `bson:"tweet" json:"tweet,omitempty"`
	Date   time.Time          `bson:"date" json:"date,omitempty"`
}

type TweetStruct struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Tweet string             `bson:"tweet" json:"tweet,omitempty"`
	Date  time.Time          `bson:"date" json:"date,omitempty"`
}

type TweetFromFollowers struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID string             `bson:"user_id" json:"user_id,omitempty"`
	Tweet  TweetStruct
}
