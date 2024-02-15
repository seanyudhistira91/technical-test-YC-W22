// utils/password.go
package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt takes a plain text password, hashes it with a random salt, and returns the hashed password.
func HashAndSalt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares a hashed password with a plain text password and returns true if they match.
func ComparePassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
