//go:build wireinject
// +build wireinject

package main

import (
	"github.com/arturdolzan/go-snapshot-server/application"
	"github.com/arturdolzan/go-snapshot-server/domain"

	"github.com/google/wire"
)

func InitializeSnapshotServerApplication() *application.SnapshotServerApplication {
	wire.Build(
		application.NewSnapshotServerApplication,
		domain.NewAwsNodeRepository,
		domain.NewAwsImageRepository,
	)
	return nil
}
