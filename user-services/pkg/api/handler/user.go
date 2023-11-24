package handler

import (
	handlerinterfaces "github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/api/handler/interfaces"
	services "github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/usecase/interfaces"
)

type UserHandler struct {
	UserUsecase services.UserUsecase
}

func NewUserHandler(Userusecase services.UserUsecase) handlerinterfaces.UserHandler {
	return &UserHandler{UserUsecase: Userusecase}
}
