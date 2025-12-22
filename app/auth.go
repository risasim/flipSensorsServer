package app

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// AuthMiddleware verifies that the sender has tha apikey
func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header"})
	}
	dotenv := goDotEnvVariable("ARDUINO_API_KEY")
	apiKey := strings.Split(authHeader, "Bearer ")[1]
	if apiKey != dotenv {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization"})
	}
	c.Next()
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
