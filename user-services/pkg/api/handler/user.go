package handler

import (
	handlerinterfaces "github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/api/handler/interfaces"
	services "github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/usecase/interfaces"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase services.UserUsecase
}

func NewUserHandler(Userusecase services.UserUsecase) handlerinterfaces.UserHandler {
	return &UserHandler{UserUsecase: Userusecase}
}

func (U *UserHandler) UserLogin(ctx *gin.Context) {
	panic("Unimplimented UserLogin --> Handler")
}

func (U *UserHandler) UserHome(ctx *gin.Context) {
	panic("Unimplimented UserHome --> Handler")
}

func (U *UserHandler) UserLogout(ctx *gin.Accounts) {
	panic("Unimplimented UserLogout --> Handler")
}
