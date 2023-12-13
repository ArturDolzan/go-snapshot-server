package repositories

import "fmt"

type AwsImageRepository struct{}

func (a AwsImageRepository) DescribeImages() {
	fmt.Println("Aws Image repo called!")
}
