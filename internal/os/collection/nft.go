package collection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type NFT struct {
	DataSource      string `json:"data_source_name" db:"data_source_name"`
	Chain           string `json:"chain" db:"chain"`
	Slug            string `json:"slug" db:"slug"`
	Name            string `json:"name" db:"name"`
	ImageURL        string `json:"large_image_url" db:"image_url"`
	DiscordURL      string `json:"discord_url" db:"discord_url"`
	URL             string `json:"external_url" db:"url"`
	SafelistStatus  string `json:"safelist_request_status" db:"safelist_status"`
	TwitterHandle   string `json:"twitter_username" db:"twitter_handle"`
	InstagramHandle string `json:"instagram_username" db:"instagram_handle"`
}

func ParseNFT(ds, chain, path string) (*NFT, error) {
	type collection struct {
		NFT NFT `json:"collection"`
	}

	result := collection{}
	result.NFT.DataSource = ds
	result.NFT.Chain = chain
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("failed to read '%s', %w", path, err)
		log.Errorf(err.Error())
		return nil, err
	}
	err = json.Unmarshal(bs, &result)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s', %w", path, err)
		log.Errorf(err.Error())
		return nil, err
	}
	return &result.NFT, nil
}
