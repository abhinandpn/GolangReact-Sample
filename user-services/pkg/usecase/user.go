package usecase

import (
	"github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/repositories/interfaces"
	services "github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/usecase/interfaces"
)

type UserUsecase struct {
	UserRepo interfaces.UserRepository
}

func NewUserUsecase(UserRepo interfaces.UserRepository) services.UserUsecase {
	return &UserUsecase{UserRepo: UserRepo}
}

func (U *UserUsecase) UserLogin() {

}

func (U *UserUsecase) UserLogout() {

}

func (U *UserUsecase) UserHome() {

}

func (U *UserUsecase) UserAccount() {

}
