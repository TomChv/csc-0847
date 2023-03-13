package constants

import "os"

const (
	UploadDir = "/tmp/uploadDir"
)

func init() {
	os.MkdirAll(UploadDir, 0755)
}
