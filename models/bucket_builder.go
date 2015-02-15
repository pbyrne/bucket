package models

import (
	"crypto/md5"
	"fmt"
	"github.com/pbyrne/bucket/util"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type BucketBuilder struct {
	bucket Bucket
	Dir    string
}

func NewBucketBuilder(bucket Bucket) BucketBuilder {
	dir, err := ioutil.TempDir("", "bucket")
	util.PanicIf(err)
	return BucketBuilder{bucket: bucket, Dir: dir}
}

func (bb BucketBuilder) Perform() {
	bb.writeIndex()
	bb.copyImages()
	bb.buildAssets()
}

func (bb BucketBuilder) Images() []Image {
	return bb.bucket.Images()
}

func (bb BucketBuilder) FingerprintedJavaScripts() []string {
	result := make([]string, 0)
	for _, jsSrc := range bb.javaScriptPaths() {
		result = append(result, bb.fingerprintedBaseName(jsSrc))
	}
	return result
}

func (bb BucketBuilder) FingerprintedStylesheets() []string {
	result := make([]string, 0)
	for _, cssSrc := range bb.stylesheetPaths() {
		result = append(result, bb.fingerprintedBaseName(cssSrc))
	}
	return result
}

func (bb BucketBuilder) writeIndex() {
	index, err := os.Create(path.Join(bb.Dir, "/index.html"))
	util.PanicIf(err)
	template, err := template.ParseFiles("templates/index.html")
	util.PanicIf(err)
	template.Execute(index, bb)
}

func (bb BucketBuilder) copyImages() {
	for _, image := range bb.bucket.Images() {
		bb.copyFile(path.Join(bb.Dir, image.BaseName()), image.Path)
	}
}

func (bb BucketBuilder) buildAssets() {
	jsDestPath := path.Join(bb.Dir, "javascripts")
	err := os.Mkdir(jsDestPath, 0744)
	util.PanicIf(err)
	for _, jsSrc := range bb.javaScriptPaths() {
		dest := path.Join(jsDestPath, bb.fingerprintedBaseName(jsSrc))
		bb.copyFile(dest, jsSrc)
	}

	cssDestPath := path.Join(bb.Dir, "stylesheets")
	err = os.Mkdir(cssDestPath, 0744)
	util.PanicIf(err)
	for _, cssSrc := range bb.stylesheetPaths() {
		dest := path.Join(cssDestPath, bb.fingerprintedBaseName(cssSrc))
		bb.copyFile(dest, cssSrc)
	}
}

func (bb BucketBuilder) copyFile(destPath string, srcPath string) {
	dest, err := os.Create(destPath)
	util.PanicIf(err)
	src, err := os.Open(srcPath)
	util.PanicIf(err)
	io.Copy(dest, src)
}

func (bb BucketBuilder) CleanUp() {
	os.RemoveAll(bb.Dir)
}

func (bb BucketBuilder) javaScriptPaths() []string {
	sourcePaths, err := filepath.Glob(filepath.Join("public/javascripts", "*"))
	util.PanicIf(err)
	return sourcePaths
}

func (bb BucketBuilder) stylesheetPaths() []string {
	sourcePaths, err := filepath.Glob(filepath.Join("public/stylesheets", "*"))
	util.PanicIf(err)
	return sourcePaths
}

func (bb BucketBuilder) fingerprintedBaseName(path string) string {
	data, err := ioutil.ReadFile(path)
	util.PanicIf(err)
	hash := md5.Sum(data)
	nameParts := strings.SplitN(filepath.Base(path), ".", 2)
	return fmt.Sprintf("%s-%x.%s", nameParts[0], hash, nameParts[1])
}
