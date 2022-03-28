// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"plan-to-remind/internal/biz"
	"plan-to-remind/internal/conf"
	"plan-to-remind/internal/data"
	"plan-to-remind/internal/server"
	"plan-to-remind/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db, cleanup := data.NewGormDb(confData)
	dataData, cleanup2, err := data.NewData(db)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	cronSpecRepo := data.NewCronSpecRepo(dataData)
	cronSpecUsecase := biz.NewCronSpecUsecase(cronSpecRepo)
	cronService := service.NewCronService(cronSpecUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, cronService, logger)
	grpcServer := server.NewGRPCServer(confServer, cronService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
