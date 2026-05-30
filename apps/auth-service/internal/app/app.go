package app

import (
	"fmt"
	"time"

	"github.com/dogukanttopcuoglu/beatflow/apps/auth-service/internal/config"
	"github.com/dogukanttopcuoglu/beatflow/apps/auth-service/internal/handler/health"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type App struct {
	cfg    config.Config
	logger *zap.Logger
	server *fiber.App
}

func New(cfg config.Config, logger *zap.Logger) *App {
	server := fiber.New(fiber.Config{
		AppName: cfg.Service.Name,
	})

	healthHandler := health.NewCheckHandler(cfg.Service.Name)
	health.RegisterRoutes(server, healthHandler)

	return &App{
		cfg:    cfg,
		logger: logger,
		server: server,
	}
}

func (a *App) Start() error {
	addr := fmt.Sprintf(":%d", a.cfg.HTTP.Port)

	a.logger.Info("starting HTTP server", zap.String("address", addr))

	return a.server.Listen(addr)
}

func (a *App) Shutdown(timeout time.Duration) error {
	a.logger.Info("shutting down HTTP server",
		zap.Duration("timeout", timeout),
	)

	return a.server.ShutdownWithTimeout(timeout)
}
