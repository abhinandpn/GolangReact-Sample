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

}

func (P *ProductUseCase) UpdateProduct() {

}

func (P *ProductUseCase) DeletProduct() {

}

// Product Listing
func (P *ProductUseCase) ViewFullProducts() {

}

func (P *ProductUseCase) ViewProductById() {

}

func (P *ProductUseCase) ViewProductByName() {

}

func (P *ProductUseCase) ViewProductsByCategory() {

}

func (P *ProductUseCase) ViewProductsBySUbCategory() {

}

func (P *ProductUseCase) ViewProductsByBrands() {

}
