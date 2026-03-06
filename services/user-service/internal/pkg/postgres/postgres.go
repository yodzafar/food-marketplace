package postgres

import (
	"fmt"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/yodzafar/food-marketpalce/user-service/config"
)

func newConnection(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DBName,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	return db, nil
}

var ProviderSet = wire.NewSet(newConnection)
