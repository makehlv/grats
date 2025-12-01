package services

import (
	"log/slog"

	"github.com/makehlv/grats/internal/builders"
	"github.com/makehlv/grats/internal/clients"
	"github.com/makehlv/grats/internal/config"
	"github.com/makehlv/grats/internal/infra/postgres"
	"github.com/makehlv/grats/internal/repositories"
	"github.com/makehlv/grats/internal/services/common"
	"github.com/makehlv/grats/internal/services/support"
	"github.com/makehlv/grats/internal/services/user"
	"github.com/makehlv/grats/internal/services/wish"
	"github.com/makehlv/grats/internal/services/wishlist"
)

type Services struct {
	User     *user.Service
	Wish     *wish.Service
	WishList *wishlist.Service
	Support  *support.Service
}

func New(
	cfg *config.Config,
	logger *slog.Logger,
	repositories *repositories.Repositories,
	clients *clients.Clients,
	builders *builders.Builders,
	db *postgres.DB,
) *Services {
	common := common.New(cfg, logger, repositories, clients, builders)
	return &Services{
		User:     user.New(cfg, logger, db, repositories, clients, builders, common),
		Wish:     wish.New(cfg, logger, db, repositories, clients, builders, common),
		WishList: wishlist.New(cfg, logger, db, repositories, clients, builders),
		Support:  support.New(cfg, logger, db, repositories, clients, builders),
	}
}
