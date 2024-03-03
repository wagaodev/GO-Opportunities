package main

import (
	"github.com/wagaodev/Go-Opportunities/config"
	"github.com/wagaodev/Go-Opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	router.Initialize()
}
