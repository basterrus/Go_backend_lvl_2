//go:generate wire
//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"strconv"

	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/lapitskyss/go_backend_2/internal/config"
	"github.com/lapitskyss/go_backend_2/internal/server"
	"github.com/lapitskyss/go_backend_2/internal/server/handler"
	"github.com/lapitskyss/go_backend_2/internal/srv/producsrv"
	"github.com/lapitskyss/go_backend_2/internal/store/elastic"
)

type App struct {
	Server *server.Server
	Log    *zap.Logger
}

var AppSet = wire.NewSet(
	InitApp,
	InitServer,
	InitContext,
	InitConfig,
	InitLogger,
	InitElastic,
)

func InitializeApp() (*App, func(), error) {
	panic(wire.Build(AppSet))
}

func InitApp(s *server.Server, log *zap.Logger) (*App, error) {
	return &App{
		Server: s,
		Log:    log,
	}, nil
}

func InitServer(conf *config.Config, es *elastic.Store, log *zap.Logger) (*server.Server, func(), error) {
	port := strconv.Itoa(conf.ServerPort)

	securitySrv := producsrv.InitProductSrv(es, log)
	h := handler.InitHandler(securitySrv, log)

	s := server.InitServer(port, h, log)

	cleanup := func() {
		_ = s.Stop()
	}

	s.Start()

	return s, cleanup, nil
}

func InitContext() (context.Context, func(), error) {
	ctx := context.Background()

	cb := func() {
		ctx.Done()
	}

	return ctx, cb, nil
}

func InitConfig() (*config.Config, error) {
	cfg := &config.Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func InitLogger() (*zap.Logger, func(), error) {
	logger, _ := zap.NewProduction()

	cleanup := func() {
		_ = logger.Sync()
	}

	return logger, cleanup, nil
}

func InitElastic(conf *config.Config) (*elastic.Store, error) {
	return elastic.InitElasticsearch(conf.ElasticsearchAddresses)
}
