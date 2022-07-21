package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type PublicMetrics struct {
	RetweetCount uint `json:"retweet_count" db:"retweet_count"`
	ReplyCount   uint `json:"reply_count" db:"reply_count"`
	LikeCount    uint `json:"like_count" db:"like_count"`
	QuoteCount   uint `json:"quote_count" db:"quote_count"`
}

type TwitterStats struct {
	Day        string `db:"day"`
	Criterion  string `db:"criterion"`
	Slug       string `db:"slug"`
	DataSource string `db:"data_source_name"`
	Followers  uint   `db:"followers"`
	SearchHits uint   `db:"search_hits"`
	PublicMetrics
}

func ParseSearchStats(chain, day, dsource, spath string) error {
	files, err := ioutil.ReadDir(spath)
	if err != nil {
		log.Error("failed to read twitter search stats files for  ", spath)
		log.Error(err)
		return err
	}
	ps := fmt.Sprintf("%c", os.PathSeparator)
	ss := strings.Split(spath, ps)
	slug := ss[len(ss)-1]
	ts := TwitterStats{Day: day, Criterion: chain, Slug: slug, DataSource: dsource}
	for _, file := range files {
		fpath := spath + ps + file.Name()
		err = parseStats(fpath, &ts)
		if err != nil {
			log.Error("failed to parse twitter search stats for ", fpath)
			log.Error(err)
		}
	}

	fmt.Printf("** %#v\n", ts)
	return nil
}

func parseStats(fpath string, stats *TwitterStats) error {
	type single struct {
		PM PublicMetrics `json:"public_metrics"`
	}
	type all struct {
		Data []single `json: "data"`
	}
	result := all{}
	bs, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Error("failed to read ", fpath)
		return err
	}
	err = json.Unmarshal(bs, &result)
	if err != nil {
		log.Error("failed to parse ", fpath)
		return err
	}
	stats.SearchHits += uint(len(result.Data))
	for _, d := range result.Data {
		stats.RetweetCount += d.PM.RetweetCount
		stats.ReplyCount += d.PM.ReplyCount
		stats.LikeCount += d.PM.LikeCount
		stats.QuoteCount += d.PM.QuoteCount
	}
	return nil
}
