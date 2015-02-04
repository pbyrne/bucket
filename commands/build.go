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
	builder := NewBucketBuilder(bucket)

	builder.Perform()
	fmt.Println(builder)
}

type BucketBuilder struct {
	bucket models.Bucket
	dir    string
}

func NewBucketBuilder(bucket models.Bucket) BucketBuilder {
	dir, err := ioutil.TempDir("", "bucket")
	util.PanicIf(err)
	return BucketBuilder{bucket: bucket, dir: dir}
}

func (bb BucketBuilder) Perform() {
	defer bb.cleanUp()

	bb.writeIndex()
	bb.copyFiles()
}

func (bb BucketBuilder) writeIndex() {
	index, err := os.Create(path.Join(bb.dir, "/index.html"))
	util.PanicIf(err)
	template, err := template.ParseFiles("templates/index.html")
	template.Execute(index, bb.bucket)
}

func (bb BucketBuilder) copyFiles() {
	for _, image := range bb.bucket.Images() {
		dest, err := os.Create(path.Join(bb.dir, image.BaseName()))
		util.PanicIf(err)
		src, err := os.Open(image.Path)
		util.PanicIf(err)
		io.Copy(dest, src)
	}
}

func (bb BucketBuilder) cleanUp() {
	os.RemoveAll(bb.dir)
}
