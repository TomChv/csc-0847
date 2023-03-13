package constants

import "os"

const (
	OriginalDir  = "/tmp/original"
	ThumbnailDir = "/tmp/thumbnail"
)

func init() {
	os.MkdirAll(OriginalDir, 0755)
	os.MkdirAll(ThumbnailDir, 0755)
}
