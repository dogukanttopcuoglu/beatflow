package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	authapp "github.com/dogukanttopcuoglu/beatflow/apps/auth-service/internal/app"
	authconfig "github.com/dogukanttopcuoglu/beatflow/apps/auth-service/internal/config"
	sharedlogger "github.com/dogukanttopcuoglu/beatflow/packages/go/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := authconfig.Load()

	if err != nil {
		panic(err)
	}

	log, err := sharedlogger.New(sharedlogger.Options{
		Environment: cfg.Service.Env,
		ServiceName: cfg.Service.Name,
		Level:       cfg.Log.Level,
	})

	if err != nil {
		panic(err)
	}

	defer log.Sync()

	log.Info("auth service starting",
		zap.String("environment", cfg.Service.Env),
		zap.Int("http_port", cfg.HTTP.Port),
	)

	application := authapp.New(cfg, log)

	serverErrors := make(chan error, 1)

	go func() {
		if err := application.Start(); err != nil {
			serverErrors <- err
		}
	}()

	shutdownSignals := make(chan os.Signal, 1)

	signal.Notify(shutdownSignals, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Error("server failed", zap.Error(err))
		os.Exit(1)
	case sig := <-shutdownSignals:
		log.Info("shutdown signal received", zap.String("signal", sig.String()))
	}

	shutdownTimeout := 5 * time.Second

	if err := application.Shutdown(shutdownTimeout); err != nil {
		log.Error("application shutdown failed", zap.Error(err))
		os.Exit(1)
	}

	log.Info("auth service stopped gracefully")

}
