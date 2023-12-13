package domain

import (
	"github.com/arturdolzan/go-snapshot-server/infrastructure/repositories"
)

type ImageRepository interface {
	DescribeImages()
}

func NewAwsImageRepository() ImageRepository {
	return &repositories.AwsImageRepository{}
}

func NewAzureImageRepository() ImageRepository {
	return &repositories.AzureImageRepository{}
}
