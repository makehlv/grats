package app

import (
	"github.com/makehlv/grats/internal/builders"
	"github.com/makehlv/grats/internal/clients"
	"github.com/makehlv/grats/internal/config"
	"github.com/makehlv/grats/internal/infra/postgres"
	"github.com/makehlv/grats/internal/infra/redis"
	"github.com/makehlv/grats/internal/repositories"
	"github.com/makehlv/grats/internal/services"
	tgbot "github.com/makehlv/tgbot"
)

func Run() {
	cfg := config.MustLoad()
	logger := MustSetupLogging(cfg)

	db := postgres.New(cfg, logger)

	redis := redis.New(cfg, logger)

	repositories := repositories.New(cfg, logger, db, redis)
	clients := clients.New(cfg, logger)
	builders := builders.New(cfg, logger)
	services := services.New(cfg, logger, repositories, clients, builders, db)

	bot := tgbot.New(&cfg.TgBot, logger, repositories.State)

	RegisterHandlers(bot, services, cfg, repositories)

	bot.Serve()
}
