package main

import (
	"github.com/arturdolzan/go-snapshot-server/application"
	"github.com/arturdolzan/go-snapshot-server/config"
	"github.com/arturdolzan/go-snapshot-server/presentation"
	"github.com/arturdolzan/go-snapshot-server/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error : %v", err)
		return
	}

	app := InitializeSnapshotServerApplication()

	application.Initialize()
	presentation.Initialize(app)

	router.Initialize()
}
