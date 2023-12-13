package routerhandler

import "github.com/arturdolzan/go-snapshot-server/config"

var (
	logger *config.Logger
)

func Initialize() {
	logger = config.GetLogger("handler")
}
