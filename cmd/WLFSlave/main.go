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
		Name:        "WLF Slave",
		Description: "Bullshit Kuber slave by Hirbod Behnam",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "master-address",
				Value: "127.0.0.1:12346",
				Usage: "Address of master to connect to",
			},
			&cli.StringFlag{
				Name:  "listen-address",
				Value: "127.0.0.1:0",
				Usage: "Address which master can connect to slave",
			},
			&cli.IntFlag{
				Name:  "max-tasks",
				Value: 4,
				Usage: "Max tasks this slave can serve",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "run the slave",
				Action: func(ctx *cli.Context) error {
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
