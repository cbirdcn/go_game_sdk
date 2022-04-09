// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"sdk/internal/biz"
	"sdk/internal/conf"
	"sdk/internal/data"
	"sdk/internal/server"
	"sdk/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewGormClient(confData, logger)
	dataData, cleanup, err := data.NewData(confData, db, logger)
	// 需要使用ent时
	// entClient := data.NewEntClient(confData, logger)
	// dataData, cleanup, err := data.NewData(confData, entClient, db, logger)
	if err != nil {
		return nil, nil, err
	}
	sdkRepo := data.NewSdkRepo(dataData, logger)
	sdkUsecase := biz.NewSdkUsecase(sdkRepo, logger)
	sdkService := service.NewSdkService(sdkUsecase)
	httpServer := server.NewHTTPServer(confServer, sdkService, logger)
	grpcServer := server.NewGRPCServer(confServer, sdkService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
