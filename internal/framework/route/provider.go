package route

import (
	"logbook-api/internal/application"
	"logbook-api/internal/config"
	"logbook-api/internal/infrastructure"
)

// Provider struct
type Provider struct {
	Config     *config.Config
	Repository *infrastructure.Repository
	Usecase    *application.Usecase
}

// New ...
func New(c *config.Config, r *infrastructure.Repository, u *application.Usecase) Provider {
	return Provider{
		Config:     c,
		Repository: r,
		Usecase:    u,
	}
}
