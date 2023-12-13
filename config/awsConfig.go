package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetAwsSession(profile string, region string) (*session.Session, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
		Config: aws.Config{
			Region: aws.String(region),
		},
	})

	if err != nil {
		logger.Errorf("Error creating AWS session %v", err.Error())
		return nil, err
	}

	return sess, nil
}
