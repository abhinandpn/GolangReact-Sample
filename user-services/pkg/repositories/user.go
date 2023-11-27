package repositories

import (
	"github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/repositories/interfaces"
	"gorm.io/gorm"
)

type UserDataBase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &UserDataBase{DB: DB}
}

func (U *UserDataBase) FindUser() {

}
