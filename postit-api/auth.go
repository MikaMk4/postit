package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func WithJWTAuth(handler gin.HandlerFunc, store Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the token from the request (Auth header)
		tokenString := GetTokenFromRequest(c)
		// validate the token
		token, err := validateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized" + err.Error()})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		// get the userID from the token
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["userId"].(string)

		_, err = store.GetUserById(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized" + err.Error()})
			c.Abort()
			return
		}
		// call the handler func and continue to the endpoint
		c.Set("userId", userID)
		handler(c)
	}
}

func GetTokenFromRequest(c *gin.Context) string {
	token := c.GetHeader("Authorization")

	if token != "" {
		return token
	}

	return ""
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(secret), nil
	})
}

func CreateJWT(id string, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    id,
		"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func ComparePasswords(hashed string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
