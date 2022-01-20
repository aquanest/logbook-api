package framework

import (
	"logbook-api/internal/application"
	"logbook-api/internal/config"
	routeProvider "logbook-api/internal/framework/route"
	"logbook-api/internal/infrastructure"
)

// Controller struct
type Controller struct {
	Config     *config.Config
	Repository *infrastructure.Repository
	Usecase    *application.Usecase
}

// NewController returns new controller
func NewController(c *config.Config, r *infrastructure.Repository, u *application.Usecase) Controller {
	return Controller{
		Config:     c,
		Repository: r,
		Usecase:    u,
	}
}

// InvokeRouteProvider ...
func (c *Controller) InvokeRouteProvider() routeProvider.Provider {
	return routeProvider.New(c.Config, c.Repository, c.Usecase)
}
