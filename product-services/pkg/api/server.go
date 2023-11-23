package http

import (
	handlerinterfaces "github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/api/handler/interfaces"
	"github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/api/routes"
	"github.com/gin-gonic/gin"
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
	// Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.ProductRoute(Engine.Group("/admin"),
		ProductHandler,
	)
	return &ServerHTTP{engine: Engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":5000")
}
