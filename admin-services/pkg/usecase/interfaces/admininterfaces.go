package interfaces

import (
	"context"
)

type AdminUseCase interface {
	AdminLogin(ctx context.Context)
	AdminHome(ctx context.Context)
	AdminLogout(ctx context.Context)
}
