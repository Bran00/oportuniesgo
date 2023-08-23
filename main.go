package main

import (
	"github.com/Bran00/oportuniesgo/config"
	"github.com/Bran00/oportuniesgo/router"
)

var (
	logger config.Logger
)

func main() {
	logger := config.GetLogger("main")
	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	//Initialize Router
	router.Initialize()
}
