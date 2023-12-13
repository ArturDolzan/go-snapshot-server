package valueObjects

import (
	"fmt"
	"strings"
)

type BlockDevice struct {
	DeviceName string `json:"deviceName"`
	VolumeId   string `json:"volumeId"`
}

func (b *BlockDevice) Validate() error {
	if !strings.HasPrefix(b.DeviceName, "/dev/") {
		return fmt.Errorf("BlockDevice name is not valid: %s", b.DeviceName)
	}

	return nil
}
