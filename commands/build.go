package commands

import (
	"fmt"
	"github.com/pbyrne/bucket/models"
	"github.com/pbyrne/bucket/util"
	"io/ioutil"
	"os"
)

func Build(bucket models.Bucket) {
	dir, err := ioutil.TempDir("", "bucket")
	util.PanicIf(err)
	index, err := os.Create(dir + "/index.html")
	util.PanicIf(err)

	index.WriteString("<h1>Hello, world!</h1>")

	fmt.Println(dir)
}
