package usecase

import (
	"context"

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
func (A *AdminUseCase) AdminLogin(ctx context.Context) {
	panic("unimplemented AdminLogin --> Handler")
}

// AdminLogin implements interfaces.AdminUseCase.
func (A *AdminUseCase) AdminHome(ctx context.Context) {
	panic("unimplemented AdminHome --> Handler")
}

// AdminLogin implements interfaces.AdminUseCase.
func (A *AdminUseCase) AdminLogout(ctx context.Context) {
	panic("unimplemented AdminLogout --> Handler")
}
