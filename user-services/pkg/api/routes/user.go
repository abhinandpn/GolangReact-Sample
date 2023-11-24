package routes

import (
	handlerinterfaces "github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/api/handler/interfaces"

	"github.com/gin-gonic/gin"
)

func UserRoute(api *gin.RouterGroup,
	UserHandler handlerinterfaces.UserHandler) {

}
