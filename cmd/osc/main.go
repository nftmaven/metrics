package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/nftmaven/metrics/internal/os/collection"
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

	var chain, criterion, day, dsource, fpath string
	app := &cli.App{
		Name:  "osc",
		Usage: "OpenSea client",
		Commands: []*cli.Command{
			{
				Name:    "process-top-100",
				Aliases: []string{"pth"},
				Usage:   "process a top-100 csv file",
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
						Usage:       "top-100 csv file path",
						Required:    true,
						Destination: &fpath,
					},
				},
				Action: func(c *cli.Context) error {
					log.Info("fpath = ", fpath)
					data, err := top100.Process(dsource, criterion, fpath)
					if err != nil {
						return err
					}
					err = top100.Persist(db, criterion, data)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "process-collection-stats",
				Aliases: []string{"pcs"},
				Usage:   "process a json file with stats for a specific collection",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "chain",
						Usage:       "the chain this collection lives on (e.g. \"ethereum\")",
						Required:    true,
						Destination: &chain,
					},
					&cli.StringFlag{
						Name:        "day",
						Usage:       "date on which the stats were captured (e.g. \"2022-06-14\")",
						Required:    true,
						Destination: &day,
					},
					&cli.StringFlag{
						Name:        "dsource",
						Usage:       "data source (e.g. \"opensea\")",
						Required:    true,
						Destination: &dsource,
					},
					&cli.StringFlag{
						Name:        "fpath",
						Usage:       "stats file path",
						Required:    true,
						Destination: &fpath,
					},
				},
				Action: func(c *cli.Context) error {
					log.Info("fpath = ", fpath)
					nft, err := collection.ParseNFT(dsource, chain, fpath)
					if err != nil {
						return err
					}
					stats, err := collection.ParseStats(dsource, chain, fpath)
					if err != nil {
						return err
					}
					err = collection.Persist(db, chain, day, nft, stats)
					if err != nil {
						return err
					}
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
