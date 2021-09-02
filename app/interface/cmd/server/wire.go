// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/biz"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/conf"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/data"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/server"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Auth, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, service.ProviderSet, data.ProviderSet, newApp))
}
