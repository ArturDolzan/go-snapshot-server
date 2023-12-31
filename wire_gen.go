// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/arturdolzan/go-snapshot-server/application"
	"github.com/arturdolzan/go-snapshot-server/domain"
)

// Injectors from wire.go:

func InitializeSnapshotServerApplication() *application.SnapshotServerApplication {
	nodeRepository := domain.NewAwsNodeRepository()
	imageRepository := domain.NewAwsImageRepository()
	snapshotServerApplication := application.NewSnapshotServerApplication(nodeRepository, imageRepository)
	return snapshotServerApplication
}
