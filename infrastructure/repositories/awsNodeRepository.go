package repositories

import (
	"github.com/arturdolzan/go-snapshot-server/config"
	"github.com/arturdolzan/go-snapshot-server/domain/entities"
	"github.com/arturdolzan/go-snapshot-server/domain/valueObjects"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AwsNodeRepository struct{}

func (a AwsNodeRepository) DescribeNodes(profile string, region string) ([]entities.Node, error) {
	session, err := config.GetAwsSession(profile, region)
	if err != nil {
		return nil, err
	}

	ec2Svc := ec2.New(session)
	input := &ec2.DescribeInstancesInput{}
	result, err := ec2Svc.DescribeInstances(input)
	if err != nil {
		return nil, err
	}

	var instances []entities.Node
	for _, reservation := range result.Reservations {
		for _, inst := range reservation.Instances {
			state := valueObjects.State{Value: aws.StringValue(inst.State.Name)}
			errState := state.Validate()
			if errState != nil {
				return nil, errState
			}

			instanceType := valueObjects.InstanceType{Value: aws.StringValue(inst.InstanceType)}
			errInstanceType := instanceType.Validate()
			if err != nil {
				return nil, errInstanceType
			}

			publicIP := ""
			if inst.PublicIpAddress != nil {
				publicIP = aws.StringValue(inst.PublicIpAddress)
			}

			privateIP := ""
			if inst.PrivateIpAddress != nil {
				privateIP = aws.StringValue(inst.PrivateIpAddress)
			}

			var blockDevices []valueObjects.BlockDevice
			for _, bdm := range inst.BlockDeviceMappings {
				blockDevice := valueObjects.BlockDevice{
					DeviceName: aws.StringValue(bdm.DeviceName),
					VolumeId:   aws.StringValue(bdm.Ebs.VolumeId),
				}
				errBlockDevice := blockDevice.Validate()
				if errBlockDevice != nil {
					return nil, errBlockDevice
				}
				blockDevices = append(blockDevices, blockDevice)
			}

			var tags []entities.Tag
			for _, tag := range inst.Tags {
				tags = append(tags, entities.Tag{
					Key:   aws.StringValue(tag.Key),
					Value: aws.StringValue(tag.Value),
				})
			}

			instanceInfo := entities.Node{
				Id:           aws.StringValue(inst.InstanceId),
				State:        state,
				InstanceType: instanceType,
				LaunchTime:   aws.TimeValue(inst.LaunchTime),
				PublicIP:     publicIP,
				PrivateIP:    privateIP,
				BlockDevices: blockDevices,
				Tags:         tags,
			}
			instances = append(instances, instanceInfo)
		}
	}

	return instances, nil
}
