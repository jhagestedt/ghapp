package main

import (
	"fmt"
	"github.com/jhagestedt/ghapp/v2/cmd"
	"github.com/urfave/cli/v2"
	"os"
	"sort"
)

var version = "dev"

func main() {
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show help",
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Show version",
	}
	app := &cli.App{
		Name:    "ghapp",
		Usage:   "GitHub app cli",
		Version: version,
		Commands: []*cli.Command{
			{
				Name:    "token",
				Usage:   "Create a GitHub app installation token",
				Aliases: []string{"t"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "GitHub app id",
						EnvVars:  []string{"GITHUB_APP_ID"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "installation-id",
						Usage:    "GitHub app installation-id",
						EnvVars:  []string{"GITHUB_APP_INSTALLATION_ID"},
						Required: true,
					},
					&cli.StringFlag{
						Name:        "private-key",
						Usage:       "GitHub app private-key",
						EnvVars:     []string{"GITHUB_APP_PRIVATE_KEY"},
						FilePath:    ".github-app-private-key.pem",
						DefaultText: ".github-app-private-key.pem",
						Required:    true,
					},
				},
				Action: func(ctx *cli.Context) error {
					id := ctx.String("id")
					installationId := ctx.String("installation-id")
					privateKey := ctx.String("private-key")
					token, err := cmd.CreateToken(id, privateKey)
					if err != nil {
						return err
					}
					installationToken, err := cmd.CreateInstallationToken(installationId, token)
					if err != nil {
						return err
					}
					fmt.Print(installationToken)
					return nil
				},
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
