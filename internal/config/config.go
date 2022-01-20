package config

import (
	"log"

	"logbook-api/internal/domain"

	"github.com/jinzhu/configor"
	"github.com/urfave/cli/v2"
)

// Config struct
type Config struct {
	WebListenAddress  string
	WebListenPort     int
	WebDebugMode      bool
	ConfigAtmosServer string
	ConfigAtmosToken  string
}

// New returns Config struct
func New(ctx *cli.Context) Config {
	cfg := Config{
		WebListenAddress:  ctx.String(domain.WebListenAddressFlagName),
		WebListenPort:     ctx.Int(domain.WebListenPortFlagName),
		WebDebugMode:      ctx.Bool(domain.WebDebugFlagName),
		ConfigAtmosServer: ctx.String(domain.ConfigAtmosServerFlagName),
		ConfigAtmosToken:  ctx.String(domain.ConfigAtmosTokenFlagName),
	}

	err := configor.New(&configor.Config{}).Load(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = checkArguments(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

func checkArguments(cfg *Config) error {
	return nil
}
