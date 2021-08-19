package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User The User Entity definition.
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	LastName  string             `bson:"last_name,omitempty" json:"last_name"`
	BirthDay  time.Time          `bson:"birth_day,omitempty" json:"birth_day"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
}
