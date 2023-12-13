package valueObjects

import "fmt"

type InstanceType struct {
	Value string `json:"instanceType"`
}

func (i *InstanceType) Validate() error {
	validInstanceTypes := []string{"t3.nano", "t3.micro", "t3.small", "t3.medium", "t3.large", "t3.xlarge", "t3.2xlarge", "m5a.2xlarge", "m5.4xlarge"}

	if !contains(validInstanceTypes, i.Value) {
		return fmt.Errorf("Instace type is not valid: %s", i.Value)
	}

	return nil
}
