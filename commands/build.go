package commands

import (
	"fmt"
	"github.com/pbyrne/bucket/models"
	"github.com/pbyrne/bucket/util"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func Build(bucket models.Bucket) {
	dir, err := ioutil.TempDir("", "bucket")
	util.PanicIf(err)

	indexPath := path.Join(dir, "/index.html")
	index, err := os.Create(indexPath)
	util.PanicIf(err)
	defer os.RemoveAll(dir)
	template, err := template.ParseFiles("templates/index.html")
	template.Execute(index, bucket)

	for _, image := range bucket.Images() {
		dest, err := os.Create(path.Join(dir, image.BaseName()))
		util.PanicIf(err)
		src, err := os.Open(image.Path)
		util.PanicIf(err)
		io.Copy(dest, src)
	}

	fmt.Println(dir)
}
