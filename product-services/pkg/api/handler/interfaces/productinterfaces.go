package interfaces

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	
	// Product
	AddProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeteleProduct(ctx *gin.Context)
	ListFullProducts(ctx *gin.Context)
	GetProductById(ctx *gin.Context)

	// Category
	// Sub-Category
	// Brand
	// Sorting
	// Searching
}
