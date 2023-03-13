package store

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"

	"github.com/TomChv/csc-847/project_2/api/constants"
)

type Client struct {
	c *firestore.Client
}

func New(ctx context.Context) (*Client, error) {
	pictureStore, err := firestore.NewClient(ctx, "cdc-847-project-2")
	if err != nil {
		return nil, err
	}

	return &Client{c: pictureStore}, nil
}

func (c *Client) Close() error {
	return c.c.Close()
}

type File struct {
	Name   string      `json:"name"`
	Url    string      `json:"url"`
	Labels interface{} `json:"labels"`
}

func (c *Client) List(ctx context.Context) ([]File, error) {
	pictures := c.c.Collection(constants.StoreCollection)

	docs, err := pictures.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var results []File
	for _, doc := range docs {
		results = append(results, File{
			Name:   doc.Ref.ID,
			Url:    fmt.Sprintf("https://storage.googleapis.com/thumbnails-cdc-847-project-2/%s", doc.Ref.ID),
			Labels: doc.Data()["labels"],
		})
	}

	return results, nil
}

func (c *Client) UpdateLabel(ctx context.Context, id string, label string) error {
	pictures := c.c.Collection(constants.StoreCollection)

	_, err := pictures.Doc(id).Update(ctx, []firestore.Update{
		{Path: "labels", Value: []string{label}},
	})

	return err
}
