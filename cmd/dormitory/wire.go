//go:build wireinject
// +build wireinject

package main

// The build tag makes sure the stub is not built in the final build.
import (
	"dormitory/internal/biz"
	"dormitory/internal/conf"
	"dormitory/internal/data"
	"dormitory/internal/server"
	"dormitory/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		newApp,
	))
}
