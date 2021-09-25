package repository

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/cjmarkham/GoBB/config"
)

func ProvideClient(config *config.PostgresqlConfig) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN("postgres://gobb:password@db.postgres:5432/gobb_dev?sslmode=disable")))

	return bun.NewDB(sqldb, pgdialect.New())
}
