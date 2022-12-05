package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	app := &cli.App{
		Name:        "WLF Client",
		Description: "Bullshit Kuber client by Hirbod Behnam",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "master-address",
				Value: "127.0.0.1:12345",
				Usage: "Address of master to connect to",
			},
			&cli.StringFlag{
				Name:  "username",
				Value: "admin",
				Usage: "Username to authorize in master",
			},
			&cli.StringFlag{
				Name:  "password",
				Value: "pass",
				Usage: "Password to authorize in master",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "job",
				Usage: "manage jobs",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add job to master",
						Flags: []cli.Flag{
							&cli.PathFlag{
								Name:  "config",
								Value: "job.json",
							},
							&cli.BoolFlag{
								Name:  "simple",
								Value: true,
							},
							&cli.PathFlag{
								Name: "executable",
							},
						},
						Action: func(ctx *cli.Context) error {
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "get jobs list",
						Action: func(ctx *cli.Context) error {
							return nil
						},
					},
					{
						Name:  "logs",
						Usage: "get logs of a job",
						Flags: []cli.Flag{},
						Action: func(ctx *cli.Context) error {
							return nil
						},
					},
				},
			},
			{
				Name:  "node",
				Usage: "manage nods",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "get a list of nodes and their status",
						Action: func(ctx *cli.Context) error {
							return nil
						},
					},
					{
						Name:  "top",
						Usage: "get node utilization",
						Action: func(ctx *cli.Context) error {
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
