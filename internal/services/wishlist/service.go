package wishlist

import (
	"log/slog"

	"github.com/makehlv/grats/internal/builders"
	"github.com/makehlv/grats/internal/clients"
	"github.com/makehlv/grats/internal/config"
	"github.com/makehlv/grats/internal/infra/postgres"
	"github.com/makehlv/grats/internal/repositories"
)

type Service struct {
	logger       *slog.Logger
	db           *postgres.DB
	repositories *repositories.Repositories
	clients      *clients.Clients
	builders     *builders.Builders
	cfg          *config.Config
}

func New(
	cfg *config.Config,
	logger *slog.Logger,
	db *postgres.DB,
	repositories *repositories.Repositories,
	clients *clients.Clients,
	builders *builders.Builders,
) *Service {
	return &Service{
		logger:       logger,
		db:           db,
		repositories: repositories,
		clients:      clients,
		builders:     builders,
		cfg:          cfg,
	}
}
