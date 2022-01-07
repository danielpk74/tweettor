package users

import (
	"github.com/danielpk74/tweettor/models"
)

func Login(email string, password string) (models.User, bool) {
	user, found, _ := UserAlreadyExists(email)
	if !found {
		return user, false
	}

	validPassword := ComparePassword([]byte(user.Password), []byte(password))
	if !validPassword {
		return user, false
	}

	return user, true
}
