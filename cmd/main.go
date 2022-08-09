package main

import (
	"flag"

	"github.com/ZeineI/corrector/internal/app"
)

func main() {
	configPath := flag.String("config-path", "./configs/config.yaml", "Path to the config file")
	flag.Parse()

	app.Run(*configPath)
}
