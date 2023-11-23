package http

import (
	handlerinterfaces "github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/api/handler/interfaces"
	"github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/api/routes"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(AdminHandler handlerinterfaces.AdminHandler) *ServerHTTP {

	Engine := gin.New()
	Engine.Use(gin.Logger())

	// For Html
	// Engine.LoadHTMLGlob("*.html")

	// For Swagger
	// Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.AdminRoute(Engine.Group("/admin"),
		AdminHandler,
	)
	return &ServerHTTP{engine: Engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":5000")
}
