package main

import (
	"github.com/pbyrne/bucket/commands"
	"github.com/pbyrne/bucket/models"
	"os"
)

func main() {
	bucket := models.NewBucket()
	cmd := parseCommand(os.Args)

	switch cmd {
	case "list":
		commands.List(bucket)
	case "version":
		commands.Version()
	default:
		commands.Usage()
	}
}

func parseCommand(args []string) string {
	if len(args) > 1 {
		return args[1]
	} else {
		return "usage"
	}
}
