package db

import (
	"context"
	"fmt"

	"github.com/TomChv/csc-0847/project_1/backend/ent"
)

type Client struct {
	*ent.Client
}

func New() (*Client, error) {
	// Load configuration
	loadConfig()

	fmt.Printf("database configuration retrieved with provider: %s on %s:%s\n", provider, host, port)

	providerFunc, exist := providers[provider]
	if !exist || providerFunc == nil {
		return nil, fmt.Errorf("provider %s does not exist", provider)
	}

	c, err := providerFunc()
	if err != nil {
		return nil, err
	}

	if err := c.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close() error {
	return c.Client.Close()
}
