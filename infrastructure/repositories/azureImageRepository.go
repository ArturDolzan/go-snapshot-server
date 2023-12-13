package repositories

import "fmt"

type AzureImageRepository struct{}

func (a AzureImageRepository) DescribeImages() {
	fmt.Println("Azure Image repo called!")
}
