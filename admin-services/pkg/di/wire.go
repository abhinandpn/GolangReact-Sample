//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/api"
	"github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/api/handler"
	"github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/config"
	"github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/database"
	"github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/repositories"
	"github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/usecase"
	"github.com/google/wire"
)

func InitializeApi(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		database.ConnectDatabase,

		//handler
		handler.NewAdminHandler,

		//usecase
		usecase.NewAdminUseCase,

		//repository
		repositories.NewAdminRepository,

		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
