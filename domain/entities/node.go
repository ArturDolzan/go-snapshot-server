package entities

import (
	"fmt"
	"time"

	"github.com/arturdolzan/go-snapshot-server/domain/valueObjects"
)

type Node struct {
	Id                        string    `json:"id"`
	LaunchTime                time.Time `json:"launchTime"`
	PublicIP                  string    `json:"publicIP"`
	PrivateIP                 string    `json:"privateIP"`
	valueObjects.State        `json:"state"`
	valueObjects.InstanceType `json:"instanceType"`
	BlockDevices              []valueObjects.BlockDevice `json:"blockDevices"`
	Tags                      []Tag                      `json:"tags"`
}

// Validations
func errModel(name string, typ string, message string) error {
	return fmt.Errorf("param: %s (type: %s) - %s", name, typ, message)
}

func (n *Node) Validate() error {
	if n.Id == "" {
		return errModel("Id", "string", "Id is required")
	}

	return nil
}
