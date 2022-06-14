package top100

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

type Top100 struct {
	Rank        int    `csv:"rank"`
	Day         string `csv:"date"`
	ID          string `csv:"id"`
	Name        string `csv:"name"`
	Slug        string `csv:"slug"`
	IsVerified  bool   `csv:"is_verified"`
	CreatedDate string `csv:"created_date"`
}

type Top100M struct {
	ID         int    `db:"id"`
	Day        string `db:"day"`
	Criterion  string `db:"criterion"`
	Rank       int    `db:"rank"`
	Slug       string `db:"slug"`
	DataSource string `db:"data_source_name"`
}

func Process(db *sqlx.DB, ds, criterion, fpath string) ([]*Top100, error) {
	data := []*Top100{}

	fh, err := os.OpenFile(fpath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		err = fmt.Errorf("failed to open file '%s', %w", fpath, err)
		log.Errorf(err.Error())
		return nil, err
	}
	err = gocsv.UnmarshalFile(fh, &data)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal '%s', %w", fpath, err)
		log.Errorf(err.Error())
		return nil, err
	}

	if len(data) == 0 {
		return data, nil
	}
	q := fmt.Sprintf(
		`INSERT INTO top100stats(day, criterion, rank, slug, data_source_name)
       VALUES ('%s', '%s', :rank, :slug, '%s')`, data[0].Day, criterion, ds)
	_, err = db.NamedExec(q, data)
	if err != nil {
		err = fmt.Errorf("failed to write to db '%s', %w", criterion, err)
		log.Errorf(err.Error())
		return nil, err
	}

	return data, nil
}
