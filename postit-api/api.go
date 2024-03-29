package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	addr  string
	store Store
}

func NewApiServer(addr string, store Store) *ApiServer {
	return &ApiServer{
		addr:  addr,
		store: store,
	}
}

func (s *ApiServer) Start() error {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.MaxAge = 12 * time.Hour
	router.Use(cors.New(config))

	subrouter := router.Group("/api/v1")

	// register services
	postsService := NewPostService(s.store)
	postsService.RegisterRoutes(subrouter)

	usersService := NewUserService(s.store)
	usersService.RegisterRoutes(subrouter)

	return router.Run(s.addr)
}
