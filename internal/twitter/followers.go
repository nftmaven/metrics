package twitter

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jmoiron/sqlx"
)

type FollowerData struct {
	Day        string `db:"day"`
	Slug       string `db:"slug"`
	Followers  uint   `db:"followers"`
	Following  uint   `db:"following"`
	TweetCount uint   `db:"tweet_count"`
	UserName   string `db:"user_name"`
}

func ParseFollowers(day, slug, fpath string) (*FollowerData, error) {
	type PM struct {
		FollowersCount uint `json:"followers_count"`
		FollowingCount uint `json:"following_count"`
		TweetCount     uint `json:"tweet_count"`
	}
	type Data struct {
		PM       PM     `json:"public_metrics"`
		UserName string `json:"username"`
	}
	type L1 struct {
		Data Data `json: "data"`
	}
	result := L1{}
	bs, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Error("failed to read ", fpath)
		return nil, err
	}
	err = json.Unmarshal(bs, &result)
	if err != nil {
		log.Error("failed to parse ", fpath)
		return nil, err
	}
	fd := FollowerData{
		Day:        day,
		Slug:       slug,
		Followers:  result.Data.PM.FollowersCount,
		Following:  result.Data.PM.FollowingCount,
		TweetCount: result.Data.PM.TweetCount,
		UserName:   result.Data.UserName,
	}

	return &fd, nil
}
func PersistFollowers(db *sqlx.DB, fd FollowerData) error {

	q := `
		INSERT INTO twitter_stats(
			day, criterion, slug, data_source_name, followers, tweet_count)
		SELECT
			:day, n.chain, :slug, n.data_source_name, :followers, :tweet_count
		FROM nft n
		WHERE n.slug = :slug AND n.twitter_handle = :user_name
			ON DUPLICATE KEY UPDATE
				followers=:followers,
				tweet_count=:tweet_count
		`
	_, err := db.NamedExec(q, fd)
	if err != nil {
		log.Error("failed to insert or update ", fd)
		log.Error(err.Error())
	}
	return nil
}
