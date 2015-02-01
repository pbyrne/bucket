package models

import (
	"github.com/pbyrne/bucket/util"
	"path/filepath"
)

type Bucket struct {
	root string
}

func NewBucket() *Bucket {
	return &Bucket{root: "/Users/pbyrne/Dropbox/Photos/Bucket"}
}

func (b Bucket) Images() []Image {
	imagePaths, err := filepath.Glob(filepath.Join(b.root, "*"))
	util.PanicIf(err)
	return ImagesFromPaths(imagePaths)
}
