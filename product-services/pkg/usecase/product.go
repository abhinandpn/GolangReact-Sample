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

// Product Basic
func (P *ProductUseCase) CreateProduct() {
	panic("unimplemented CreateProduct --> Usecase")
}

func (P *ProductUseCase) UpdateProduct() {
	panic("unimplemented CreateProduct --> Usecase")
}

func (P *ProductUseCase) DeletProduct() {
	panic("unimplemented CreateProduct --> Usecase")
}

// Product Listing
func (P *ProductUseCase) ViewFullProducts() {
	panic("unimplemented CreateProduct --> Usecase")
}

func (P *ProductUseCase) ViewProductById() {
	panic("unimplemented CreateProduct --> Usecase")
}

func (P *ProductUseCase) ViewProductByName() {
	panic("unimplemented CreateProduct --> Usecase")
}

func (P *ProductUseCase) ViewProductsByCategory() {
	panic("unimplemented CreateProduct --> Usecase")
}

func (P *ProductUseCase) ViewProductsBySUbCategory() {
	panic("unimplemented CreateProduct --> Usecase")
}

func (P *ProductUseCase) ViewProductsByBrands() {
	panic("unimplemented CreateProduct --> Usecase")
}
