package valueObjects

import "fmt"

type State struct {
	Value string `json:"state"`
}

func (s *State) Validate() error {
	validStates := []string{"pending", "running", "shutting-down", "terminated", "stopping", "stopped"}

	if !contains(validStates, s.Value) {
		return fmt.Errorf("State is not valid: %s", s.Value)
	}

	return nil
}
