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
	panic("unimplemented AddProduct --> Repositories")
}

func (P *ProductDatabase) UpdateProduct() {
	panic("unimplemented UpdateProduct --> Repositories")
}

func (P *ProductDatabase) DeleteProduct() {
	panic("unimplemented DeleteProduct --> Repositories")
}

// Product filtering

func (P *ProductDatabase) GetFullProducts() {
	panic("unimplemented GetFullProducts --> Repositories")
}

func (P *ProductDatabase) GetProductById() {
	panic("unimplemented GetProductById --> Repositories")
}

func (P *ProductDatabase) GetProductByName() {
	panic("unimplemented GetProductByName --> Repositories")
}

func (P *ProductDatabase) GetProductByCategory() {
	panic("unimplemented GetProductByCategory --> Repositories")
}

func (P *ProductDatabase) GetProductByBrand() {
	panic("unimplemented GetProductByBrand --> Repositories")
}
