package repositories

import (
	"github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/repositories/interfaces"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB: DB}
}
