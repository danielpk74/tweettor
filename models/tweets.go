package models

import "time"

type Tweet struct {
	UserID string    `bson:"user_id" json:"user_id,omitempty"`
	Tweet  string    `bson:"tweet" json:"tweet,omitempty"`
	Date   time.Time `bson:"date" json:"date,omitempty"`
}
