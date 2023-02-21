package db

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"net"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"

	"github.com/TomChv/csc-0847/project_1/backend/ent"
	"github.com/TomChv/csc-0847/project_1/backend/utils"
)

// gcloud implements ProviderFunc to create a
// client to a gcloud sql instance.
func gcloud() (*Client, error) {
	instanceName := utils.ForceGetEnv("DB_INSTANCE_NAME")

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		return nil, err
	}

	config, err := pgx.ParseConfig(fmt.Sprintf("user=%s password=%s database=%s", user, password, name))
	if err != nil {
		return nil, err
	}

	config.DialFunc = func(ctx context.Context, network string, instance string) (net.Conn, error) {
		return d.Dial(ctx, instanceName)
	}

	dbURI := stdlib.RegisterConnConfig(config)

	sqlClient, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, err
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, sqlClient)
	client := ent.NewClient(ent.Driver(drv))

	return &Client{client}, nil
}
