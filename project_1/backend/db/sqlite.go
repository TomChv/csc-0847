package db

import (
	"fmt"
	"github.com/google/uuid"

	"github.com/TomChv/csc-0847/project_1/backend/ent"

	_ "github.com/mattn/go-sqlite3"
)

// sqlite implements ProviderFunc to create a
// client to a sqlite database.
// Only use this on test purpose.
func sqlite() (*Client, error) {
	client, err := ent.Open("sqlite3", fmt.Sprintf("/tmp/%s:ent?mode=memory&cache=shared&_fk=1", uuid.New()))

	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}
