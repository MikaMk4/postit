package main

import (
	"fmt"
	"net/http"
	"strconv"

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
	router.GET("/board/:bid/posts", s.getPostsByBoardId)
	router.GET("/board/:bid/posts/:pid", s.getPost)
	router.POST("/board/:bid/posts", WithJWTAuth(s.createPost, s.store))
	router.PUT("/board/:bid/posts/:pid", WithJWTAuth(s.updatePost, s.store))
	router.DELETE("/board/:bid/posts/:pid", WithJWTAuth(s.deletePost, s.store))
	router.GET("/user/:uid/posts", s.getPostsByUserId)
}

func (s *PostService) getPostsByBoardId(c *gin.Context) {
	bid, err := strconv.Atoi(c.Param("bid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "board id must be an integer"})
		return
	}

	posts, err := s.store.GetPostsByBoardId(int64(bid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (s *PostService) getPostsByUserId(c *gin.Context) {
	uid := c.Param("uid")

	posts, err := s.store.GetPostsByUserId(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (s *PostService) getPost(c *gin.Context) {
	pid, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "post id must be an integer"})
		return
	}

	post, err := s.store.GetPost(int64(pid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (s *PostService) createPost(c *gin.Context) {
	bid, err := strconv.Atoi(c.Param("bid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "board id must be an integer"})
		return
	}

	var post Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validatePostPayload(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.AuthorID = c.GetString("userId")
	post.BoardID = int64(bid)

	u, err := s.store.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, u)
}

func (s *PostService) updatePost(c *gin.Context) {
	var post Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validatePostPayload(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := s.store.UpdatePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (s *PostService) deletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
		return
	}

	err = s.store.DeletePost(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
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
