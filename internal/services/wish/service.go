package wish

import (
	"context"
	"log/slog"

	"github.com/makehlv/grats/internal/builders"
	"github.com/makehlv/grats/internal/clients"
	"github.com/makehlv/grats/internal/config"
	"github.com/makehlv/grats/internal/infra/postgres"
	"github.com/makehlv/grats/internal/repositories"
	tgbot "github.com/makehlv/tgbot"
)

type UserRegistration interface {
	RegisterOrUpdateUser(ctx context.Context, scope *tgbot.Scope) error
}

type Service struct {
	userRegistration UserRegistration
	logger           *slog.Logger
	db               *postgres.DB
	repositories     *repositories.Repositories
	clients          *clients.Clients
	builders         *builders.Builders
	cfg              *config.Config
}

func New(
	cfg *config.Config,
	logger *slog.Logger,
	db *postgres.DB,
	repositories *repositories.Repositories,
	clients *clients.Clients,
	builders *builders.Builders,
	userRegistration UserRegistration,
) *Service {
	return &Service{
		userRegistration: userRegistration,
		logger:           logger,
		db:               db,
		repositories:     repositories,
		clients:          clients,
		builders:         builders,
		cfg:              cfg,
	}
}
