package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

	dsn := getDSN()
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var dsource, criterion, fpath string
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
						Name:        "dsource",
						Usage:       "data source (e.g. \"opensea\")",
						Required:    true,
						Destination: &dsource,
					},
					&cli.StringFlag{
						Name:        "criterion",
						Usage:       "either one of the chains (e.g. \"ethereum\") or \"global\"",
						Required:    true,
						Destination: &criterion,
					},
					&cli.StringFlag{
						Name:        "fpath",
						Usage:       "top-100 file path",
						Required:    true,
						Destination: &fpath,
					},
				},
				Action: func(c *cli.Context) error {
					log.Info("fpath = ", fpath)
					data, err := top100.Process(db, dsource, criterion, fpath)
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

func getDSN() string {
	var (
		host, port, user, passwd, database string
		present                            bool
	)

	host, present = os.LookupEnv("NFTMAVEN_DB_HOST")
	if !present {
		log.Fatal("NFTMAVEN_DB_HOST variable not set")
	}
	port, present = os.LookupEnv("NFTMAVEN_DB_PORT")
	if !present {
		log.Fatal("NFTMAVEN_DB_PORT variable not set")
	}
	user, present = os.LookupEnv("NFTMAVEN_DB_USER")
	if !present {
		log.Fatal("NFTMAVEN_DB_USER variable not set")
	}
	passwd, present = os.LookupEnv("NFTMAVEN_DB_PASSWORD")
	if !present {
		log.Fatal("NFTMAVEN_DB_PASSWORD variable not set")
	}
	database, present = os.LookupEnv("NFTMAVEN_DB_DATABASE")
	if !present {
		log.Fatal("NFTMAVEN_DB_DATABASE variable not set")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true&time_zone=UTC", user, passwd, host, port, database)
	return dsn
}
