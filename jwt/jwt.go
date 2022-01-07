package jwt

import (
	"github.com/danielpk74/tweettor/models"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWT(u models.User) (string, error) {
	secretKey := []byte("ThisIsTheMostSecuredKey_group")

	payload := jwtgo.MapClaims{
		"email":     u.Email,
		"name":      u.Name,
		"last_name": u.LastName,
		"birth_day": u.BirthDay,
		"biography": u.Biography,
		"_id":       u.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
