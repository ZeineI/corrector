package main

import (
	"fmt"
)

func main() {
	cfg, err := config.newConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cfg)
}
