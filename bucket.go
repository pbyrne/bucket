package main

import (
	"fmt"
	"github.com/pbyrne/bucket/image"
	"github.com/pbyrne/bucket/util"
	"path/filepath"
)

func main() {
	root := filepath.Clean("/Users/pbyrne/Dropbox/Photos/Bucket")
	imagePaths, err := filepath.Glob(filepath.Join(root, "*"))
	util.PanicIf(err)
	images := image.ImagesFromPaths(imagePaths)

	fmt.Println("Scanning", root)
	fmt.Println(images[0])
	fmt.Println(images[0].Size())
}
