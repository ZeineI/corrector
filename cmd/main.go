package main

import (
	"fmt"

	"github.com/ZeineI/corrector/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cfg)
}
