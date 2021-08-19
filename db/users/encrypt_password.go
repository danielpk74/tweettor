package users

import "golang.org/x/crypto/bcrypt"

const cost = 8

func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
