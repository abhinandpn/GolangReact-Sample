package http

import (
	handlerinterfaces "github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/api/handler/interfaces"
	"github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/api/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(ProductHandler handlerinterfaces.ProductHandler) *ServerHTTP {

	Engine := gin.New()
	Engine.Use(gin.Logger())

	// For Html
	// Engine.LoadHTMLGlob("*.html")

	// For Swagger
	Engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.ProductRoute(Engine.Group("/product"),
		ProductHandler,
	)
	return &ServerHTTP{engine: Engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":5005")
}
