package interfaces

import (
	"context"

	"github.com/abhinandpn/MicroServices-GoLang/admin-services/pkg/models"
)

type AdminUseCase interface {
	AdminLogin(ctx context.Context, admin models.Admin) (models.Admin, error)
	// AdminHome(ctx context.Context)
	// AdminLogout()
}
