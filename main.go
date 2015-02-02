package main

import (
	"github.com/pbyrne/bucket/commands"
	"github.com/pbyrne/bucket/models"
)

func main() {
	bucket := models.NewBucket()

	commands.List(bucket)
}
