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
				Usage:   "Create a GitHub App installation token",
				Aliases: []string{"t"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "GitHub App id",
						EnvVars:  []string{"GHAPP_ID", "GITHUB_APP_ID"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "install-id",
						Usage:    "GitHub App installation id",
						EnvVars:  []string{"GHAPP_INSTALL_ID", "GITHUB_APP_INSTALL_ID"},
						Required: true,
					},
					&cli.StringFlag{
						Name:    "private-key",
						Usage:   "GitHub App private key",
						EnvVars: []string{"GHAPP_PRIVATE_KEY", "GITHUB_APP_PRIVATE_KEY"},
					},
					&cli.StringFlag{
						Name:    "private-key-file",
						Usage:   "GitHub App private key file like .ghapp-private-key.pem",
						EnvVars: []string{"GHAPP_PRIVATE_KEY_FILE", "GITHUB_APP_PRIVATE_KEY_FILE"},
					},
				},
				Action: func(ctx *cli.Context) error {
					if !ctx.IsSet("private-key") {
						if !ctx.IsSet("private-key-file") {
							return fmt.Errorf("private-key or private-key-file not set")
						}
						privateKeyFile := ctx.String("private-key-file")
						privateKeyBytes, err := os.ReadFile(privateKeyFile)
						if err != nil {
							return err
						}
						err = ctx.Set("private-key", string(privateKeyBytes))
						if err != nil {
							return err
						}
					}
					id := ctx.String("id")
					installId := ctx.String("install-id")
					privateKey := ctx.String("private-key")
					token, err := cmd.CreateToken(id, privateKey)
					if err != nil {
						return err
					}
					installationToken, err := cmd.CreateInstallationToken(installId, token)
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
	err := app.Run(os.Args)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
