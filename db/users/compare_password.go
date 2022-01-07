package users

import "golang.org/x/crypto/bcrypt"

// ComparePassword Compare a Hash password with a string.
func ComparePassword(hashPassword []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashPassword, password)
	if err != nil {
		return false
	}
	return true
}
