package repositories

import (
	"github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/repositories/interfaces"
	"gorm.io/gorm"
)

type ProductDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepositorie {
	return &ProductDatabase{DB: DB}
}

// Product Basic

func (P *ProductDatabase) AddProduct() {

}

func (P *ProductDatabase) UpdateProduct() {

}

func (P *ProductDatabase) DeleteProduct() {

}

// Product filtering

func (P *ProductDatabase) GetFullProducts() {

}

func (P *ProductDatabase) GetProductById() {

}

func (P *ProductDatabase) GetProductByName() {

}

func (P *ProductDatabase) GetProductByCategory() {

}

func (P *ProductDatabase) GetProductByBrand() {

}
