package usecase

import (
	"context"

	"github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/models"
	"github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/repositories/interfaces"
	services "github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/usecase/interfaces"
)
 
type AdminUseCase struct {
	AdminRepo interfaces.AdminRepository
}

func NewAdminUseCase(adminRepo interfaces.AdminRepository) services.AdminUseCase {
	return &AdminUseCase{AdminRepo: adminRepo}
}

// AdminLogin implements interfaces.AdminUseCase.
func (*AdminUseCase) AdminLogin(ctx context.Context, admin models.Admin) (models.Admin, error) {
	panic("unimplemented")
}
