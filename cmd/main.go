package main

import (
	"flag"

	"github.com/ZeineI/corrector/internal/app"
)

func main() {
	configPath := flag.String("config-path", "./configs/config.yaml", "Path to the config file")
	flag.Parse()

	app.Run(*configPath)

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// <-quit

	// fmt.Println("\napp shutting down")
}
