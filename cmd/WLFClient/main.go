package main

import (
	"WLF/internal/client"
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
							&cli.Uint64Flag{
								Name:  "needed_memory",
								Usage: "needed memory to run this job in bytes",
							},
							&cli.Uint64Flag{
								Name:  "needed_cores",
								Usage: "needed cores to run this job",
							},
							&cli.Uint64Flag{
								Name:  "needed_disk",
								Usage: "needed disk to run this job in bytes",
							},
						},
						Action: func(ctx *cli.Context) error {
							var neededCores *uint32
							needCores64 := nilOnZero(ctx, "needed_cores")
							if needCores64 != nil {
								neededCores = new(uint32)
								*neededCores = uint32(*needCores64)
							}
							return createMasterAndAuth(ctx).AddJob(
								ctx.String("config"),
								ctx.String("executable"),
								nilOnZero(ctx, "needed_memory"),
								nilOnZero(ctx, "needed_disk"),
								neededCores,
							)
						},
					},
					{
						Name:  "list",
						Usage: "get jobs list",
						Action: func(ctx *cli.Context) error {
							return createMasterAndAuth(ctx).PrintJobList()
						},
					},
					{
						Name:  "logs",
						Usage: "get logs of a job",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "id",
								Usage:    "The job ID to get it's logs",
								Required: true,
							},
							&cli.Uint64Flag{
								Name:  "count",
								Value: 5,
								Usage: "Number of lines to print",
							},
							&cli.BoolFlag{
								Name:  "live",
								Value: false,
								Usage: "Show live logs",
							},
							&cli.BoolFlag{
								Name:  "head",
								Value: false,
								Usage: "Show logs from top",
							},
							&cli.BoolFlag{
								Name:  "stderr",
								Value: false,
								Usage: "Show logs of stderr instead of stdout",
							},
						},
						Action: func(ctx *cli.Context) error {
							return createMasterAndAuth(ctx).LogsOfJob(
								ctx.String("id"),
								uint32(ctx.Uint64("count")),
								ctx.Bool("live"),
								!ctx.Bool("head"),
								ctx.Bool("stderr"),
							)
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
							return createMasterAndAuth(ctx).PrintNodeList()
						},
					},
					{
						Name:  "top",
						Usage: "get node utilization",
						Action: func(ctx *cli.Context) error {
							return createMasterAndAuth(ctx).PrintNodeTop()
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

// createMasterAndAuth will connect to master and send the authorization message
func createMasterAndAuth(ctx *cli.Context) *client.MasterSettings {
	m := &client.MasterSettings{
		Address:  ctx.String("master-address"),
		Username: ctx.String("username"),
		Password: ctx.String("password"),
	}
	err := m.Auth()
	if err != nil {
		m.Close()
		log.WithError(err).WithField("settings", *m).Fatalln("cannot connect to master")
	}
	return m
}

func nilOnZero(ctx *cli.Context, name string) *uint64 {
	flag := ctx.Uint64(name)
	if flag == 0 {
		return nil
	}
	return &flag
}
