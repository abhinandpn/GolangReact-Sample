package handler

import (
	handlerinterfaces "github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/api/handler/interfaces"
	services "github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/usecase/interfaces"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminUseCase services.AdminUseCase
}

func NewAdminHandler(adminUseCase services.AdminUseCase) handlerinterfaces.AdminHandler {
	return &AdminHandler{AdminUseCase: adminUseCase}
}

// AdminHome implements interfaces.AdminHandler.
func (*AdminHandler) AdminHome(ctx *gin.Context) {
	panic("unimplemented")
}

// AdminLogin implements interfaces.AdminHandler.
func (*AdminHandler) AdminLogin(ctx *gin.Context) {
	panic("unimplemented")
}

// AdminLogout implements interfaces.AdminHandler.
func (*AdminHandler) AdminLogout(ctx *gin.Context) {
	panic("unimplemented")
}
