package db

import (
	"fmt"

	"github.com/TomChv/csc-0847/project_1/backend/ent"

	_ "github.com/lib/pq"
)

// local implements ProviderFunc to create a
// client to a common postgresql connection.
func local() (*Client, error) {
	client, err := ent.Open("postgres", fmt.Sprintf(
		"host=%s "+
			"port=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"sslmode=disable", host, port, user, password, name))

	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}
