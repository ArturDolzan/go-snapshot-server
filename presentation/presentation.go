package presentation

import (
	"github.com/arturdolzan/go-snapshot-server/application"
	"github.com/arturdolzan/go-snapshot-server/config"
)

var (
	logger *config.Logger
	app    *application.SnapshotServerApplication
)

func Initialize(application *application.SnapshotServerApplication) {
	logger = config.GetLogger("presentation")
	app = application
}
