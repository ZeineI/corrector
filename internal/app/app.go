package app

import (
	"fmt"
	"log"

	"github.com/ZeineI/corrector/config"
	"github.com/ZeineI/corrector/internal/api"
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
	fmt.Println(resp)

	// server := server.NewServer(cfg)

	// if err := server.Run(cfg, logger); err != nil {
	// 	logger.Debug(err)
	// 	return
	// }
}
