package cli

import (
	"logbook-api/internal/domain"

	"github.com/urfave/cli/v2"
)

// registerWebListenAddressFlag
func registerWebListenAddressFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    domain.WebListenAddressFlagName,
			Usage:   domain.WebListenAddressFlagUsage,
			Aliases: domain.WebListenAddressFlagAlias,
			Value:   domain.WebListenAddressFlagDefaultValue,
		},
	}
}

// registerWebListenPortFlag
func registerWebListenPortFlag() []cli.Flag {
	return []cli.Flag{
		&cli.IntFlag{
			Name:    domain.WebListenPortFlagName,
			Usage:   domain.WebListenPortFlagUsage,
			Aliases: domain.WebListenPortFlagAlias,
			Value:   domain.WebListenPortFlagDefaultValue,
		},
	}
}

// registerWebDebugFlag
func registerWebDebugFlag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    domain.WebDebugFlagName,
			Usage:   domain.WebDebugFlagUsage,
			Aliases: domain.WebDebugFlagAlias,
			Value:   domain.WebDebugFlagDefaultValue,
		},
	}
}

// registerConfigAtmosServerFlag
func registerConfigAtmosServerFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     domain.ConfigAtmosServerFlagName,
			Usage:    domain.ConfigAtmosServerFlagUsage,
			Aliases:  domain.ConfigAtmosServerFlagAlias,
			EnvVars:  domain.ConfigAtmosServerFlagEnvVars,
			Required: domain.ConfigAtmosServerFlagRequired,
		},
	}
}

// registerConfigAtmosTokenFlag
func registerConfigAtmosTokenFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     domain.ConfigAtmosTokenFlagName,
			Usage:    domain.ConfigAtmosTokenFlagUsage,
			Aliases:  domain.ConfigAtmosTokenFlagAlias,
			EnvVars:  domain.ConfigAtmosTokenFlagEnvVars,
			Required: domain.ConfigAtmosTokenFlagRequired,
		},
	}
}
