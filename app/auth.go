package app

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// AuthMiddleware verifies that the sender has tha apikey
func AuthMiddleware() gin.HandlerFunc {
	// Need to check the api key from the .env file
	return func(c *gin.Context) {}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
