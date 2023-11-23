package handler

import (
	handlerinterfaces "github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/api/handler/interfaces"

	services "github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/usecase/interfaces"
)

type ProductHandler struct {
	ProductUseCase services.ProductUseCase
}

func NewProductHandler(productUsecase services.ProductUseCase) handlerinterfaces.ProductHandler {
	return &ProductHandler{ProductUseCase: productUsecase}
}
