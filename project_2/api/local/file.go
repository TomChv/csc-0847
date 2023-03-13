package local

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/TomChv/csc-847/project_2/api/constants"
)

func WriteFile(filename string, file multipart.File) error {
	filePath := filepath.Join(constants.UploadDir, filename)
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFile(filename string) {
	filePath := filepath.Join(constants.UploadDir, filename)

	os.Remove(filePath)
}
