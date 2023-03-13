package pictures

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/TomChv/csc-847/project_2/api/constants"
	"github.com/TomChv/csc-847/project_2/api/local"
	"github.com/TomChv/csc-847/project_2/api/s3"
	"github.com/TomChv/csc-847/project_2/api/utils"
)

func List(c *gin.Context) {
	fmt.Println("List - Create s3Client")
	s3Client, err := s3.New(c, constants.UploadBucket)
	if err != nil {
		utils.NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("List - List s3 files")
	files, err := s3Client.List(c)
	if err != nil {
		utils.NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("List - Return response")
	c.JSON(http.StatusOK, map[string]any{
		"files_number": len(files),
		"files":        files,
	})
}

func Add(c *gin.Context) {
	var metadata s3.Metadata

	fmt.Println("Add - Parse file")
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		utils.NewHTTPError(c, http.StatusBadRequest, err)
		return
	}

	fmt.Println("Add - Write file to local directory")
	filename := header.Filename
	if err := local.WriteFile(filename, file); err != nil {
		utils.NewHTTPError(c, http.StatusBadRequest, err)
		return
	}

	fmt.Println("Add - Retrieve metadata")
	rawMetadata := c.Request.FormValue("metadata")
	if rawMetadata == "" {
		utils.NewHTTPError(c, http.StatusBadRequest, ErrMissingMetadata)
		return
	}

	fmt.Println("Add - Unmarshal metadata")
	if err := json.Unmarshal([]byte(rawMetadata), &metadata); err != nil {
		utils.NewHTTPError(c, http.StatusBadRequest, err)
		return
	}

	fmt.Println("Add - Create s3 client")
	s3Client, err := s3.New(c, constants.UploadBucket)
	if err != nil {
		utils.NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("Add - Upload file to s3")
	err = s3Client.Upload(c, filename, metadata)
	if err != nil {
		utils.NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("Add - Return response")
	c.JSON(http.StatusCreated, map[string]string{
		"message": fmt.Sprintf("file %s successfully created", filename),
	})
}

func Update(c *gin.Context) {
	fmt.Println("Update - Retrieve name in url param")
	filename := c.Param("name")
	if filename == "" {
		utils.NewHTTPError(c, http.StatusBadRequest, ErrMissingMetadata)
		return
	}

	var metadata struct {
		Author   string `json:"author,omitempty"`
		Date     string `json:"date,omitempty"`
		Location string `json:"location,omitempty"`
		Label    string `json:"label,omitempty"`
	}

	fmt.Println("Update - Retrieve metadata from form value")
	rawMetadata := c.Request.FormValue("metadata")
	if rawMetadata == "" {
		utils.NewHTTPError(c, http.StatusBadRequest, ErrMissingMetadata)
		return
	}

	fmt.Println("Update - Unmarshal metadata")
	if err := json.Unmarshal([]byte(rawMetadata), &metadata); err != nil {
		utils.NewHTTPError(c, http.StatusBadRequest, err)
		return
	}

	fmt.Println("Update - Create s3 client")
	s3Client, err := s3.New(c, constants.UploadBucket)
	if err != nil {
		utils.NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("Update - Update metadata in s3")
	if err := s3Client.UpdateMetadata(c, filename, s3.Metadata{
		Author:   metadata.Author,
		Date:     metadata.Date,
		Location: metadata.Location,
		Label:    metadata.Label,
	}); err != nil {
		utils.NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("Update - Return response")
	c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("filename %s successfully updated", filename),
	})
}

func Delete(c *gin.Context) {
	fmt.Println("Delete - Retrieve name in url param")
	filename := c.Param("name")
	if filename == "" {
		utils.NewHTTPError(c, http.StatusBadRequest, ErrMissingMetadata)
		return
	}

	fmt.Println("Delete - Create s3 client")
	s3Client, err := s3.New(c, constants.UploadBucket)
	if err != nil {
		utils.NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("Delete - Delete file from s3")
	if err := s3Client.Delete(c, filename); err != nil {
		utils.NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("Delete - Return response")
	c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("file %s successfully deleted", filename),
	})
}
