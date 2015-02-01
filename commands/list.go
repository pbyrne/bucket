package commands

import (
	"fmt"
	"github.com/pbyrne/bucket/models"
)

func List(bucket models.Bucket) {
	for _, image := range bucket.Images() {
		fmt.Println(image.BaseName(), image.Size(), "bytes")
	}
}
