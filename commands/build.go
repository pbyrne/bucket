package commands

import (
	"fmt"
	"github.com/pbyrne/bucket/models"
)

func Build(bucket models.Bucket) {
	builder := models.NewBucketBuilder(bucket)

	builder.Perform()
	fmt.Println("Built to", builder.Dir)
}
