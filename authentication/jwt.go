package authentication

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secretKey []byte

func getSecret() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	secret := os.Getenv("JWT_SECRET")
	secretKey = []byte(secret)
	return nil
}

func getRole(username string) string {
	if username == "senior" {
		return "senior"
	}
	return "employee"
}

func CreateToken(username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,                         // Subject (user identifier)
		"iss": "todo-app",                       // Issuer
		"aud": getRole(username),                // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	fmt.Printf("Token claims added: %+v\n", claims)
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check for verification errors
	if err != nil {
		return err
	}

	// Check if the token is valid
	if !token.Valid {
		return err
	}

	// Return the verified token
	return nil
}
