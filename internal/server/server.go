package server

import (
	"github.com/ZeineI/corrector/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (server *Server) Run(cfg *config.Config, logger *zap.SugaredLogger) error {
	server.router.GET("/", text)
	return server.router.Run()
}

func text(c *gin.Context) {
	c.String(200, "Hello")
}
