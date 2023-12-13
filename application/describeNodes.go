package application

import (
	"github.com/arturdolzan/go-snapshot-server/domain/entities"
)

func (a *SnapshotServerApplication) DescribeNodes(profile string, region string) ([]entities.Node, error) {
	nodes, err := a.RepoNode.DescribeNodes(profile, region)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}
