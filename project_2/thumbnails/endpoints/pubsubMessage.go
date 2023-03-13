package endpoints

import (
	"cloud.google.com/go/pubsub"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/TomChv/csc-847/project_2/thumbnails/constants"
	"github.com/TomChv/csc-847/project_2/thumbnails/image"
	"github.com/TomChv/csc-847/project_2/thumbnails/s3"
)

func PubSubMessage(c *gin.Context) {
	var message struct {
		Message pubsub.Message
	}

	err := c.BindJSON(&message)
	if err != nil {
		NewHTTPError(c, http.StatusBadRequest, err)
		return
	}

	var object struct {
		Bucket string `json:"bucket"`
		Name   string `json:"name"`
	}

	err = json.Unmarshal(message.Message.Data, &object)
	if err != nil {
		NewHTTPError(c, http.StatusBadRequest, err)
		return
	}

	s3Client, err := s3.New(c)
	if err != nil {
		NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	originalFile := filepath.Join(constants.OriginalDir, object.Name)
	metadata, err := s3Client.Download(c, object.Bucket, object.Name, originalFile)
	if err != nil {
		NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	defer os.Remove(originalFile)
	log.Printf("Download picture %s/%s into %s", object.Bucket, object.Name, originalFile)

	thumbnail := image.NewThumbnail()
	defer thumbnail.Terminate()

	if err := thumbnail.Read(originalFile); err != nil {
		NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	if err := thumbnail.Resize(400, 400); err != nil {
		NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	thumbfile := filepath.Join(constants.ThumbnailDir, object.Name)
	defer os.Remove(thumbfile)

	if err := thumbnail.Write(thumbfile); err != nil {
		NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	log.Printf("Create local thumbnail %s into %s", object.Name, thumbfile)

	if err := s3Client.Upload(c, thumbfile, "thumbnails-cdc-847-project-2", object.Name, metadata); err != nil {
		NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	log.Printf("Upload thumbnail %s to gs://%s/%s", thumbfile, "thumbnails-cdc-847-project-2", object.Name)

	c.Status(http.StatusOK)
}
