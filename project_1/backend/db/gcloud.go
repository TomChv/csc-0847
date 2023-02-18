package db

import "fmt"

// gcloud implements ProviderFunc to create a
// client to a gcloud sql instance.
func gcloud() (*Client, error) {
	return nil, fmt.Errorf("not implemented")
}
