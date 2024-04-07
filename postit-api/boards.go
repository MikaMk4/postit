package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type BoardService struct {
	store Store
}

func NewBoardService(store Store) *BoardService {
	return &BoardService{
		store: store,
	}
}

func (s *BoardService) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/boards", s.getBoards)
	router.POST("/boards", WithJWTAuth(s.createBoard, s.store))
	router.GET("/boards/:id", s.getBoard)
	router.PUT("/boards/:id", WithJWTAuth(s.updateBoard, s.store))
	router.DELETE("/boards/:id", WithJWTAuth(s.deleteBoard, s.store))
}

func (s *BoardService) getBoards(c *gin.Context) {
	boards, err := s.store.GetBoards()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, boards)
}

func (s *BoardService) getBoard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	board, err := s.store.GetBoard(int64(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, board)
}

func (s *BoardService) createBoard(c *gin.Context) {
	var board Board
	if err := c.BindJSON(&board); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	u, err := s.store.CreateBoard(&board)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, u)
}

func (s *BoardService) updateBoard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var board Board
	if err := c.BindJSON(&board); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	b, err := s.store.GetBoard(int64(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if b.CreatorId != c.GetString("userId") {
		c.JSON(403, gin.H{"error": "forbidden"})
		return
	}

	board.ID = int64(id)
	u, err := s.store.UpdateBoard(&board)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, u)
}

func (s *BoardService) deleteBoard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	board, err := s.store.GetBoard(int64(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if board.CreatorId != c.GetString("userId") {
		c.JSON(403, gin.H{"error": "forbidden"})
		return
	}

	err = s.store.DeleteBoard(int64(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}
