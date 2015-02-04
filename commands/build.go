package commands

import (
	"fmt"
	"github.com/pbyrne/bucket/models"
	"github.com/pbyrne/bucket/util"
	"html/template"
	"io/ioutil"
	"os"
)

func Build(bucket models.Bucket) {
	dir, err := ioutil.TempDir("", "bucket")
	util.PanicIf(err)
	indexPath := dir + "/index.html"
	index, err := os.Create(indexPath)
	util.PanicIf(err)
	defer os.RemoveAll(dir)
	template, err := template.ParseFiles("templates/index.html")

	template.Execute(index, bucket)

	data, err := ioutil.ReadFile(indexPath)
	util.PanicIf(err)
	fmt.Println(string(data))
}
