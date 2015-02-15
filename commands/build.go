package commands

import (
	"fmt"
	"github.com/pbyrne/bucket/models"
)

func Build(bucket models.Bucket) {
	builder := models.NewBucketBuilder(bucket)
	defer builder.CleanUp()

	builder.Perform()
	fmt.Println(builder)
}
