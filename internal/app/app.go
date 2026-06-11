package app

import (
	"context"
	"database/sql"

	"github.com/Mariano-SI/twitter-api/internal/config"
	"github.com/Mariano-SI/twitter-api/internal/infra/storage"
	r2storage "github.com/Mariano-SI/twitter-api/internal/infra/storage/r2"
	"github.com/Mariano-SI/twitter-api/pkg/internalSql"
	"github.com/Mariano-SI/twitter-api/pkg/r2"
)

type Deps struct {
	Config     *config.Config
	DB         *sql.DB
	Storage    storage.Storage
	Transactor internalSql.Transactor
}

func New() (*Deps, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	db, err := internalSql.ConnectMySQL(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	if err != nil {
		return nil, err
	}

	r2Client, err := r2.NewClient(
		context.Background(),
		cfg.R2AccountID,
		cfg.R2AccessKeyID,
		cfg.R2SecretAccessKey,
	)
	if err != nil {
		db.Close()
		return nil, err
	}

	return &Deps{
		Config:     cfg,
		DB:         db,
		Storage:    r2storage.NewStorage(r2Client, cfg.R2Bucket, cfg.R2PublicURL),
		Transactor: internalSql.NewTransactor(db),
	}, nil
}
