package usecase

import (
	"github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/repositories/interfaces"
	services "github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/usecase/interfaces"
)

type ProductUseCase struct {
	ProductRepo interfaces.ProductRepositorie
}

func NewProductUsecase(productRepo interfaces.ProductRepositorie) services.ProductUseCase {
	return &ProductUseCase{ProductRepo: productRepo}
}
