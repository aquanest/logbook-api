package application

import (
	"logbook-api/internal/config"
	"logbook-api/internal/infrastructure"

	adminUsecase "logbook-api/internal/application/admin"
	divelogUsecase "logbook-api/internal/application/divelog"
)

// Usecase struct
type Usecase struct {
	Config     *config.Config
	Repository *infrastructure.Repository
}

// New returns Usecase struct
func New(c *config.Config, r *infrastructure.Repository) Usecase {
	return Usecase{
		Config:     c,
		Repository: r,
	}
}

// InvokeAdminUsecase ...
func (u *Usecase) InvokeAdminUsecase() *adminUsecase.Usecase {
	return &adminUsecase.Usecase{
		Config:     u.Config,
		Repository: u.Repository,
	}
}

// InvokeDivelogUsecase ...
func (u *Usecase) InvokeDivelogUsecase() *divelogUsecase.Usecase {
	return &divelogUsecase.Usecase{
		Config:     u.Config,
		Repository: u.Repository,
	}
}
