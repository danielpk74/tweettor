package routers

import (
	"errors"
	"github.com/danielpk74/tweettor/db/users"
	"github.com/danielpk74/tweettor/models"
	jwtgo "github.com/dgrijalva/jwt-go"
	"strings"
)

/*
Email
IDUser
*/
var Email string
var IDUser string

//ProcessToken process a Token to get all values
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	secretKey := []byte("ThisIsTheMostSecuredKey_group")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwtgo.ParseWithClaims(token, claims, func(tk *jwtgo.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	_, found, ID := users.UserAlreadyExists(claims.Email)
	if found {
		Email = claims.Email
		IDUser = claims.ID.Hex()

		return claims, found, ID, nil
	}

	return claims, false, string(""), errors.New("invalid token")
}
