package repository

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/cjmarkham/GoBB/config"
)

func ProvideClient(config *config.PostgresqlConfig) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.DSN)))

	bundebug.NewQueryHook(
		bundebug.FromEnv("BUNDEBUG"),
	)

	return bun.NewDB(sqldb, pgdialect.New())
}
