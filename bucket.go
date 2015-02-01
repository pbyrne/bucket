package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Image struct {
	path string
}

func (i Image) Size() int64 {
	file, err := os.Open(i.path)
	panicIf(err)
	stat, err := file.Stat()
	panicIf(err)
	return stat.Size()
}

func ImagesFromPaths(ps []string) []Image {
	var result []Image

	for _, path := range ps {
		result = append(result, Image{path: path})
	}

	return result
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	root := filepath.Clean("/Users/pbyrne/Dropbox/Photos/Bucket")
	imagePaths, err := filepath.Glob(filepath.Join(root, "*"))
	panicIf(err)
	images := ImagesFromPaths(imagePaths)

	fmt.Println("Scanning", root)
	fmt.Println(images[0])
	fmt.Println(images[0].Size())
}
