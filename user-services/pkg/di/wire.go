//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/api"
	"github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/api/handler"
	"github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/database"
	"github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/repositories"
	"github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/usecase"
	"github.com/google/wire"

	"github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/config"
)

func InitializeApi(cfg config.Config) (*http.ServeHTTP, error) {
	wire.Build(
		database.ConnectDatabase,

		//handler
		handler.NewUserHandler,

		//usecase
		usecase.NewUserUsecase,

		//respository
		repositories.NewUserRepository,

		http.NewServeHttp,
	)

	return &http.ServeHTTP{}, nil
}
