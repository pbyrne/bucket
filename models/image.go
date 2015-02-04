package models

import (
	"github.com/pbyrne/bucket/util"
	"os"
	"path/filepath"
	"time"
)

type Image struct {
	Path string
}

func (i Image) Size() int64 {
	return i.stats().Size()
}

func (i Image) ModTime() time.Time {
	return i.stats().ModTime()
}

func (i Image) BaseName() string {
	return filepath.Base(i.Path)
}

func (i Image) stats() os.FileInfo {
	file, err := os.Open(i.Path)
	util.PanicIf(err)
	stat, err := file.Stat()
	util.PanicIf(err)
	return stat
}

func ImagesFromPaths(ps []string) []Image {
	var result []Image

	for _, path := range ps {
		result = append(result, Image{Path: path})
	}

	return result
}
