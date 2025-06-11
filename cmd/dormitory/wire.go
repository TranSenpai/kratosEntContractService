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

func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		// newApp,
		newAppWithData,
	))
}
