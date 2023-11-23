//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/api"
	"github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/api/handler"
	"github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/config"
	"github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/database"
	"github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/repositories"
	"github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/usecase"
	"github.com/google/wire"
)

func InitializeApi(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		database.ConnectDatabase,

		//handler
		handler.NewProductHandler,

		//usecase
		usecase.NewProductUsecase,

		//repository
		repositories.NewProductRepository,

		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
