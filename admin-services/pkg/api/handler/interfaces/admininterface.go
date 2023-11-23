package interfaces

import "github.com/gin-gonic/gin"

type AdminHandler interface {
	AdminLogin(ctx *gin.Context)
	AdminHome(ctx *gin.Context)
	AdminLogout(ctx *gin.Context)
}
