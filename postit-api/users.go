package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	store Store
}

func NewUserService(store Store) *UserService {
	return &UserService{
		store: store,
	}
}

func (s *UserService) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/signup", s.registerUser)
	router.POST("/login", s.loginUser)
}

func (s *UserService) registerUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateUserPayload(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = hashedPassword

	u, err := s.store.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := createAndSetAuthCookie(c, u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func (s *UserService) loginUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateUserPayload(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := (s.store.GetUserByName(user.Username))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if u == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	if !ComparePasswords(u.Password, []byte(user.Password)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	token, err := createAndSetAuthCookie(c, u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func validateUserPayload(user *User) error {

	if user.Username == "" {
		return fmt.Errorf("username is required")
	}

	if user.Password == "" {
		return fmt.Errorf("password is required")
	}

	return nil
}

func createAndSetAuthCookie(c *gin.Context, id int64) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	token, err := CreateJWT(id, secret)
	if err != nil {
		return "", err
	}

	c.SetCookie("Authorization", token, 60*60*24, "/", "", false, true)

	return token, nil
}
