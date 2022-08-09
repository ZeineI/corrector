package server

import (
	"net/http"

	"github.com/ZeineI/corrector/config"
	"go.uber.org/zap"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: cfg.Server.Port,
		},
	}
}

func (server *Server) Run(cfg *config.Config, logger *zap.SugaredLogger) error {
	return nil
}
