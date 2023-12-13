package dtos

import "fmt"

type ProfileRegionDto struct {
	Profile string `json:"profile"`
	Region  string `json:"region"`
}

func errModel(name string, typ string, message string) error {
	return fmt.Errorf("param: %s (type: %s) - %s", name, typ, message)
}

func (p *ProfileRegionDto) Validate() error {
	if p.Profile == "" {
		return errModel("profile", "string", "profile is required")
	}

	if p.Region == "" {
		return errModel("region", "string", "region is required")
	}

	return nil
}
