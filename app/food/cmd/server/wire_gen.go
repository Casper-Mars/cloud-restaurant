// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/Casper-Mars/cloud-restaurant/app/food/internal/biz"
	"github.com/Casper-Mars/cloud-restaurant/app/food/internal/conf"
	"github.com/Casper-Mars/cloud-restaurant/app/food/internal/data"
	"github.com/Casper-Mars/cloud-restaurant/app/food/internal/server"
	"github.com/Casper-Mars/cloud-restaurant/app/food/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	foodRepo := data.NewFoodRepo(dataData, logger)
	foodUsecase := biz.NewFoodUsecase(foodRepo, logger)
	foodService := service.NewFoodService(foodUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, foodService)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
