package image

import (
	"github.com/pbyrne/bucket/util"
	"os"
)

type Image struct {
	path string
}

func (i Image) Size() int64 {
	file, err := os.Open(i.path)
	util.PanicIf(err)
	stat, err := file.Stat()
	util.PanicIf(err)
	return stat.Size()
}

func ImagesFromPaths(ps []string) []Image {
	var result []Image

	for _, path := range ps {
		result = append(result, Image{path: path})
	}

	return result
}
