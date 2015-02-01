package main

import (
	"fmt"
	"github.com/pbyrne/bucket/models"
)

func main() {
	bucket := models.NewBucket()

	image := bucket.Images()[0]
	fmt.Println(image)
	fmt.Println(image.Size())
}
