package main

import (
	"fmt"
	"log"

	"github.com/ZeineI/corrector/config"
	logger "github.com/ZeineI/corrector/pkg/log"
)

func main() {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Printf("Logger initialization error: %v", err)
		return
	}
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error(err)
		return
	}
	fmt.Println(cfg, logger)
}
