package models

type Relationship struct {
	UserId         string `bson:"user_id" json:"user_id"`
	FollowedUserId string `bson:"followed_user_id" json:"followed_user_id"`
	Status         bool   `json:"status,omitempty"`
}
