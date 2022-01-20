package infrastructure

import (
	"logbook-api/internal/config"

	atmosRepository "logbook-api/internal/infrastructure/atmos"
)

// Repository struct
type Repository struct {
	Config *config.Config
}

// New returns Repository struct
func New(c *config.Config) Repository {
	return Repository{
		Config: c,
	}
}

// InvokeAtmosRepository returns new AtmosRepository
func (r *Repository) InvokeAtmosRepository() *atmosRepository.Repository {
	return &atmosRepository.Repository{
		Config: r.Config,
	}
}
