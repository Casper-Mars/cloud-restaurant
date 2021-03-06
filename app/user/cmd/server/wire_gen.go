// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/Casper-Mars/cloud-restaurant/app/user/internal/biz"
	"github.com/Casper-Mars/cloud-restaurant/app/user/internal/conf"
	"github.com/Casper-Mars/cloud-restaurant/app/user/internal/data"
	"github.com/Casper-Mars/cloud-restaurant/app/user/internal/server"
	"github.com/Casper-Mars/cloud-restaurant/app/user/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	userService := service.NewUserService(userUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, userService)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
