package top100

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
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
