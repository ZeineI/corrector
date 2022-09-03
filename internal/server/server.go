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

func (server *Server) Run(cfg *config.Config, logger *zap.SugaredLogger, text []string) error {
	server.router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"texts": text,
		})
	})
	return server.router.Run()
}
