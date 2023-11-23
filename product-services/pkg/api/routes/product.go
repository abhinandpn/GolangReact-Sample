package routes

import (
	handlerinterfaces "github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/api/handler/interfaces"

	"github.com/gin-gonic/gin"
)

func ProductRoute(api *gin.RouterGroup,
	ProductHandler handlerinterfaces.ProductHandler) {

}
