package wish

import (
	"log/slog"

	"github.com/makehlv/grats/internal/config"
	"github.com/makehlv/grats/internal/infra/postgres"
)

type Repository struct {
	logger *slog.Logger
	cfg    *config.Config
	db     *postgres.DB
}

func New(cfg *config.Config, logger *slog.Logger, db *postgres.DB) *Repository {
	return &Repository{
		logger: logger,
		cfg:    cfg,
		db:     db,
	}
}
