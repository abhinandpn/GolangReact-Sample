package repositories

import (
	"github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/repositories/interfaces"
	"gorm.io/gorm"
)

type ProductDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepositorie {
	return ProductDatabase{DB: DB}
}
