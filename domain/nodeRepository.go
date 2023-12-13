package domain

import (
	"github.com/arturdolzan/go-snapshot-server/domain/entities"
	"github.com/arturdolzan/go-snapshot-server/infrastructure/repositories"
)

type NodeRepository interface {
	DescribeNodes(profile string, region string) ([]entities.Node, error)
}

func NewAwsNodeRepository() NodeRepository {
	return &repositories.AwsNodeRepository{}
}
