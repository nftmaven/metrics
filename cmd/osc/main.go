package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/nftmaven/metrics/internal/os/top100"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	bts, rev, version string
	log               = logrus.New()
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}
	version = fmt.Sprintf("%s::%s", bts, rev)
	log.Info("version = ", version)
	var fpath string
	app := &cli.App{
		Name:  "osc",
		Usage: "OpenSea client",
		Commands: []*cli.Command{
			{
				Name:    "process-top-100",
				Aliases: []string{"pth"},
				Usage:   "process top-100 files",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "fpath",
						Usage:       "top-100 file path",
						Required:    true,
						Destination: &fpath,
					},
				},
				Action: func(c *cli.Context) error {
					log.Info("fpath = ", fpath)
					data, err := top100.Process(fpath)
					if err != nil {
						return err
					}
					log.Infof("%v", data)
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
