package main

import (
	"github.com/IgorVianadF/goLang/config"
	"github.com/IgorVianadF/goLang/router"
	"github.com/joho/godotenv"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	logger.Info("Loading environments...")

	if err := godotenv.Load(); err != nil {
		logger.Errorf("Error creating environments: %v", err)
	}

	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.Initialize()
}
