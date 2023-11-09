package http

import (
	"fmt"

	"github.com/erbalo/hexago/internal/adapters/http/router"
	"github.com/erbalo/hexago/internal/app/card"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cardService card.Service
	engine      *gin.Engine
}

func NewServer(cardService card.Service) *Server {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	return &Server{
		cardService,
		engine,
	}
}

func (server *Server) ConfigureRoutes() {
	api := server.engine.Group("/v1")

	cardRouter := router.NewCardEntry(server.cardService)
	cardRouter.HandleRoutes(api)
}

func (server *Server) Run() {
	fmt.Println("Running server...")
	server.engine.Run(":8080")
}
