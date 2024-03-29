package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostService struct {
	store Store
}

func NewPostService(store Store) *PostService {
	return &PostService{
		store: store,
	}
}

func (s *PostService) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/posts", WithJWTAuth(s.getPosts, s.store))
	router.POST("/posts", WithJWTAuth(s.createPost, s.store))
}

func (s *PostService) getPosts(c *gin.Context) {
	posts, err := s.store.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (s *PostService) createPost(c *gin.Context) {
	var post Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validatePostPayload(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := s.store.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func validatePostPayload(post *Post) error {
	if post.Title == "" {
		return fmt.Errorf("title is required")
	}

	if post.Content == "" {
		return fmt.Errorf("content is required")
	}

	return nil
}
