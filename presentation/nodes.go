package presentation

import (
	"github.com/arturdolzan/go-snapshot-server/domain/entities"
)

func DescribeNodes(profile string, region string) ([]entities.Node, error) {

	nodes, err := app.DescribeNodes(profile, region)
	if err != nil {
		logger.Errorf("Describe nodes error: %v", err.Error())
		return nil, err
	}

	return nodes, nil
}
