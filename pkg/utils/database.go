package utils

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"runtime"
	"subscription_tracker/pkg/config"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewDatabase(*bun.DB) (*bun.DB, error) {
	cfg := config.New()
	databaseCfg := cfg.Database
	maxOpenConns := 4 * runtime.GOMAXPROCS(0)

	opts := []pgdriver.Option{
		pgdriver.WithAddr(fmt.Sprintf("%s:%d", databaseCfg.Host, databaseCfg.Port)),
		pgdriver.WithDatabase(databaseCfg.Name),
		pgdriver.WithUser(databaseCfg.Username),
		pgdriver.WithPassword(databaseCfg.Password),
	}

	if databaseCfg.SSLMode {
		// TODO: Revisit the guardrails ignore.
		opts = append(opts, pgdriver.WithTLSConfig(&tls.Config{ServerName: databaseCfg.Host})) // guardrails-disable-line
	} else {
		opts = append(opts, pgdriver.WithInsecure(true))
	}

	pgconn := pgdriver.NewConnector(opts...)
	sqldb := sql.OpenDB(pgconn)
	sqldb.SetMaxOpenConns(maxOpenConns)
	sqldb.SetMaxIdleConns(maxOpenConns)

	db := bun.NewDB(sqldb, pgdialect.New())

	_, err := db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	return db, nil

}
