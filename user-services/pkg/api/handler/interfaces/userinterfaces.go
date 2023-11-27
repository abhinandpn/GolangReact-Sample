package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {
	// UserBasic
	UserLogin(ctx *gin.Context)
	UserHome(ctx *gin.Context)
	UserLogout(ctx *gin.Accounts)
}
