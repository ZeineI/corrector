package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ZeineI/corrector/config"
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

	server := server.NewServer(cfg)

	if err := server.Run(cfg, logger); err != nil {
		logger.Debug(err)
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logger.Debug("app shutting down")
}
