package main

import "github.com/gin-gonic/gin"

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
	subrouter := router.Group("/api/v1")

	// register services
	postsService := NewPostService(s.store)
	postsService.RegisterRoutes(subrouter)

	usersService := NewUserService(s.store)
	usersService.RegisterRoutes(subrouter)

	return router.Run(s.addr)
}
