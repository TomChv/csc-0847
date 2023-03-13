package s3

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"os"
	"path/filepath"

	"github.com/TomChv/csc-847/project_2/api/constants"
)

type Client struct {
	c      *storage.Client
	bucket string
}

func New(ctx context.Context, bucket string) (*Client, error) {
	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{c: c, bucket: bucket}, nil
}

func (c *Client) Delete(ctx context.Context, name string) error {
	return c.c.Bucket(c.bucket).Object(name).Delete(ctx)
}

type File struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Metadata
}

func (c *Client) List(ctx context.Context) ([]File, error) {
	var files []File

	it := c.c.Bucket(c.bucket).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("Bucket(%q).Objects: %v", c.bucket, err)
		}

		files = append(files, File{
			Name: attrs.Name,
			URL:  fmt.Sprintf("%s/%s/%s", constants.ProjectURL, constants.ResultBucket, attrs.Name),
			Metadata: Metadata{
				Author:   attrs.Metadata["Author"],
				Date:     attrs.Metadata["Date"],
				Location: attrs.Metadata["Location"],
				Label:    attrs.Metadata["Label"],
			},
		})
	}

	return files, nil
}

func (c *Client) UpdateMetadata(ctx context.Context, name string, metadata Metadata) error {
	o := c.c.Bucket(c.bucket).Object(name)
	attrs, err := o.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("object.Attrs: %v", err)
	}

	o = o.If(storage.Conditions{MetagenerationMatch: attrs.Metageneration})

	newMetadata := metadata
	if newMetadata.Author == "" {
		newMetadata.Author = attrs.Metadata["Author"]
	}

	if newMetadata.Date == "" {
		newMetadata.Date = attrs.Metadata["Date"]
	}

	if newMetadata.Location == "" {
		newMetadata.Location = attrs.Metadata["Location"]
	}

	if newMetadata.Label == "" {
		newMetadata.Label = attrs.Metadata["Label"]
	}

	// Update metadata
	attrsUpdateRequest := storage.ObjectAttrsToUpdate{
		Metadata: map[string]string{
			"Author":   newMetadata.Author,
			"Location": newMetadata.Location,
			"Date":     newMetadata.Date,
			"Label":    newMetadata.Label,
		},
	}

	if _, err := o.Update(ctx, attrsUpdateRequest); err != nil {
		return err
	}

	return nil
}

func (c *Client) Upload(ctx context.Context, name string, metadata Metadata) error {
	filename := filepath.Join(constants.UploadDir, name)

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	dest := c.c.Bucket(c.bucket).Object(name).NewWriter(ctx)

	_, err = io.Copy(dest, file)
	if err != nil {
		dest.Close()
		return err
	}
	dest.Close()

	return c.UpdateMetadata(ctx, name, metadata)
}
