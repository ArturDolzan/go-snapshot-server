package application

import (
	"github.com/arturdolzan/go-snapshot-server/config"
	"github.com/arturdolzan/go-snapshot-server/domain"
)

var (
	logger *config.Logger
)

func Initialize() {
	logger = config.GetLogger("application")
}

type SnapshotServerApplication struct {
	RepoNode  domain.NodeRepository
	RepoImage domain.ImageRepository
}

func NewSnapshotServerApplication(repoNode domain.NodeRepository, repoImage domain.ImageRepository) *SnapshotServerApplication {
	return &SnapshotServerApplication{RepoNode: repoNode, RepoImage: repoImage}
}
