package http

import (
	handlerinterfaces "github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/api/handler/interfaces"
	"github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/api/routes"

	"github.com/gin-gonic/gin"
)

type ServeHTTP struct {
	engine *gin.Engine
}

func NewServeHttp(UserHandler handlerinterfaces.UserHandler) *ServeHTTP {
	Engine := gin.New()
	Engine.Use(gin.Logger())

	// For Html
	// Engine.LoadHTMLGlob("*.html")

	// For Swagger
	// Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.UserRoute(Engine.Group("/user"),
		UserHandler,
	)
	return &ServeHTTP{engine: Engine}
}

func (sh *ServeHTTP) Start() {
	sh.engine.Run("5010")
}
