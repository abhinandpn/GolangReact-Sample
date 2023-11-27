package routes

import (
	handlerinterfaces "github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/api/handler/interfaces"

	"github.com/gin-gonic/gin"
)

func ProductRoute(api *gin.RouterGroup,
	ProductHandler handlerinterfaces.ProductHandler) {

	product := api.Group("/product")
	{
		// Product Basic
		product.POST("/add", ProductHandler.AddProduct)      // Product Adding
		product.PATCH("/:id", ProductHandler.UpdateProduct)  // Product Update
		product.DELETE("/:id", ProductHandler.DeteleProduct) // Product Delete
		product.GET("/id", ProductHandler.GetProductById)    // Product GetById

		// Listing
		product.GET("/all", ProductHandler.ListFullProducts) // Full Products Listing
		product.GET("/category/:name")                       // Products Listing By category4
		product.GET("/brand/:name")                          // Producr Listign By name
		product.GET("/subcategory/:name")                    // Product Listign By sub-category
	}

}
