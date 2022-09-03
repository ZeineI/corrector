package app

import (
	"log"

	"github.com/ZeineI/corrector/config"
	"github.com/ZeineI/corrector/internal/api"
	"github.com/ZeineI/corrector/internal/server"
	logger "github.com/ZeineI/corrector/pkg/log"
)

func Run(configPath string) {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("Logger initialization error: %v", err)
	}
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal(err)
	}

	resp, err := api.GetResponse(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	router := server.NewServer()

	if err := router.Run(cfg, logger, resp); err != nil {
		logger.Info(err)
		return
	}
}
