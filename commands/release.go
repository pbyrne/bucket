package commands

import (
	"fmt"
	"github.com/pbyrne/bucket/models"
	"os"
	"os/exec"
)

func Release(bucket models.Bucket) {
	builder := models.NewBucketBuilder(bucket)
	defer builder.CleanUp()

	builder.Perform()
	fmt.Println("Built to", builder.Dir)
	rsyncBucket(builder.Dir, "bucket.patrickbyrne.net:/var/www/bucketbeta.patrickbyrne.net/public")
}

func rsyncBucket(src, dest string) {
	srcWithSlash := fmt.Sprintf("%s/", src)
	cmd := exec.Command("rsync", "-avz", "--delete", srcWithSlash, dest)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
