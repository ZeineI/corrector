package main

import (
	"corrector/config"
	"fmt"
)

func main() {
	configViper := config.ConfigNew()
	comf, err := config.ParseConfig(configViper)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(comf.Server.Host)
	fmt.Println(comf.Server.Port)
}
