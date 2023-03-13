package s3

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"os"
)

type Client struct {
	c *storage.Client
}

func New(ctx context.Context) (*Client, error) {
	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{c}, nil
}

func (c *Client) Download(ctx context.Context, bucket, name, dest string) (map[string]string, error) {
	file, err := c.c.Bucket(bucket).Object(name).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return nil, err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		return nil, err
	}

	attrs, err := c.c.Bucket(bucket).Object(name).Attrs(ctx)
	if err != nil {
		return nil, err
	}

	return attrs.Metadata, nil
}

func (c *Client) Upload(ctx context.Context, filename, bucket, name string, metadata map[string]string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	dest := c.c.Bucket(bucket).Object(name).NewWriter(ctx)

	_, err = io.Copy(dest, file)
	if err != nil {
		dest.Close()
		return err
	}
	dest.Close()

	o := c.c.Bucket(bucket).Object(name)
	attrs, err := o.Attrs(ctx)
	if err != nil {
		fmt.Println(fmt.Errorf("object.Attrs: %v", err))
		return fmt.Errorf("object.Attrs: %v", err)
	}

	o = o.If(storage.Conditions{MetagenerationMatch: attrs.Metageneration})

	// Update metadata
	attrsUpdateRequest := storage.ObjectAttrsToUpdate{
		Metadata: metadata,
	}

	if _, err := o.Update(ctx, attrsUpdateRequest); err != nil {
		fmt.Println(fmt.Errorf("failed to update metadata: %v"), err)
		return err
	}

	fmt.Println("metadata updated!")
	return nil
}
