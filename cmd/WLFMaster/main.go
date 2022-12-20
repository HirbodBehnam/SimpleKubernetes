package main

import (
	"WLF/internal/master"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"math/rand"
	"os"
	"time"
)

func main() {
	log.SetLevel(log.TraceLevel)
	rand.Seed(time.Now().Unix())
	app := &cli.App{
		Name:        "WLF Master",
		Description: "Bullshit Kuber master by Hirbod Behnam",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "client-listen-address",
				Value: "127.0.0.1:12345",
				Usage: "Address which clients connect to give commands",
			},
			&cli.StringFlag{
				Name:  "slave-listen-address",
				Value: "127.0.0.1:12346",
				Usage: "Address which slaves connect",
			},
			&cli.StringFlag{
				Name:  "username",
				Value: "admin",
				Usage: "Username which is authorized on master",
			},
			&cli.StringFlag{
				Name:  "password",
				Value: "pass",
				Usage: "Password which is authorized on master",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "run the master",
				Action: runMaster,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func runMaster(ctx *cli.Context) error {
	server := master.Server{
		Users: map[string]string{
			ctx.String("username"): ctx.String("password"),
		},
		Salves: master.NewSlaveList(),
	}
	server.RunServer(ctx.String("client-listen-address"), ctx.String("slave-listen-address"))
	return nil
}
