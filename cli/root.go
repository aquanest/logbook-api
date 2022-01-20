package cli

import (
	"fmt"
	"logbook-api/internal/application"
	"logbook-api/internal/config"
	"logbook-api/internal/domain"
	"logbook-api/internal/framework"
	"logbook-api/internal/infrastructure"
	"os"

	"github.com/urfave/cli/v2"
)

// Preset parameters for this command
const (
	cmdName      string = "logbook"
	cmdUsage     string = "Serve your dive-logs"
	cmdUsageText string = "logbook COMMAND [options...]"
	cmdVersion   string = "0.1.0"
)

// registerGlobalFlags returns global flags
func registerGlobalFlags() []cli.Flag {
	flags := []cli.Flag{}
	flags = append(flags, registerWebListenAddressFlag()...)
	flags = append(flags, registerWebListenPortFlag()...)
	flags = append(flags, registerWebDebugFlag()...)
	flags = append(flags, registerConfigAtmosServerFlag()...)
	flags = append(flags, registerConfigAtmosTokenFlag()...)
	return flags
}

// Start is a entrypoint of this command
func Start() {
	cmd := &cli.App{
		Name:      cmdName,
		HelpName:  cmdName,
		Usage:     cmdUsage,
		UsageText: cmdUsageText,
		Version:   cmdVersion,
		Flags:     registerGlobalFlags(),
		Action: func(ctx *cli.Context) error {
			c := config.New(ctx)
			r := infrastructure.New(&c)
			u := application.New(&c, &r)
			server := framework.New(&c, &r, &u)

			server.Boot(
				ctx.String(domain.WebListenAddressFlagName),
				ctx.Int(domain.WebListenPortFlagName),
				ctx.Bool(domain.WebDebugFlagName),
			)
			return nil
		},
	}

	err := cmd.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
}
