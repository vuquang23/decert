package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"decert/internal/pkg/builder"
	"decert/internal/pkg/components"
	"decert/internal/pkg/config"
	"decert/internal/pkg/migrations"
)

func main() {
	flagUp := "up"
	flagDown := "down"
	defaultMigrationDir := "migrations/decert"

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "api",
				Aliases: []string{"a"},
				Usage:   "Run api server",
				Action: func(ctx *cli.Context) error {
					if err := config.Load(""); err != nil {
						return err
					}
					if err := components.Init(); err != nil {
						return err
					}
					server, _ := builder.NewApiServer()
					server.Run(config.Instance().Http.BindAddress)
					return nil
				},
			},
			{
				Name:    "migrate",
				Aliases: []string{"m"},
				Usage:   "Run db migration",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  flagUp,
						Value: -1,
					},
					&cli.IntFlag{
						Name:  flagDown,
						Value: -1,
					},
				},
				Action: func(ctx *cli.Context) error {
					err := config.Load("")
					if err != nil {
						return err
					}
					up := ctx.Int(flagUp)
					down := ctx.Int(flagDown)
					if up == -1 && down == -1 {
						fmt.Println("No up or down migration declared")
						return nil
					}
					if up != -1 && down != -1 {
						return errors.New("[ERROR] Both up and down migration declared. Stop the migration")
					}
					m, err := migrations.NewMigration(defaultMigrationDir)
					if err != nil {
						fmt.Println("Can not create migration " + err.Error())
					}
					if up != -1 {
						return m.MigrateUp(up)
					} else {
						return m.MigrateDown(down)
					}
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
