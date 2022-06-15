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
	ID          int    `csv:"-" db:"id"`
	Rank        int    `csv:"rank" db:"rank"`
	Day         string `csv:"date" db:"day"`
	DataSource  string `csv:"-" db:"data_source_name"`
	Name        string `csv:"name"`
	Slug        string `csv:"slug" db:"slug"`
	IsVerified  bool   `csv:"is_verified"`
	CreatedDate string `csv:"created_date"`
}

func Process(ds, criterion, fpath string) ([]*Top100, error) {
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

	return data, nil
}

func Persist(db *sqlx.DB, criterion string, data []*Top100) error {
	if len(data) == 0 {
		return nil
	}
	q := fmt.Sprintf(
		`INSERT INTO top100stats(day, criterion, rank, slug, data_source_name)
		 VALUES ('%s', '%s', :rank, :slug, '%s')`, data[0].Day, criterion,
		data[0].DataSource)
	_, err := db.NamedExec(q, data)
	if err != nil {
		err = fmt.Errorf("failed to write to db '%s', %w", criterion, err)
		log.Errorf(err.Error())
		return err
	}

	return nil
}
