package collection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/shopspring/decimal"
)

type Stats struct {
	Day1Volume    decimal.Decimal `json:"one_day_volume"`
	Day1Change    decimal.Decimal `json:"one_day_change"`
	Day1Sales     decimal.Decimal `json:"one_day_sales"`
	Day1AvgPrice  decimal.Decimal `json:"one_day_average_price"`
	Day7Volume    decimal.Decimal `json:"seven_day_volume"`
	Day7Change    decimal.Decimal `json:"seven_day_change"`
	Day7Sales     decimal.Decimal `json:"seven_day_sales"`
	Day7AvgPrice  decimal.Decimal `json:"seven_day_average_price"`
	Day30Volume   decimal.Decimal `json:"thirty_day_volume"`
	Day30Change   decimal.Decimal `json:"thirty_day_change"`
	Day30Sales    decimal.Decimal `json:"thirty_day_sales"`
	Day30AvgPrice decimal.Decimal `json:"thirty_day_average_price"`
	TotalVolume   decimal.Decimal `json:"total_volume"`
	TotalSales    decimal.Decimal `json:"total_sales"`
	TotalSupply   decimal.Decimal `json:"total_supply"`
	Owners        decimal.Decimal `json:"num_owners"`
	AvgPrice      decimal.Decimal `json:"average_price"`
	MarketCap     decimal.Decimal `json:"market_cap"`
	FloorPrice    decimal.Decimal `json:"floor_price"`
}

func ParseStats(ds, chain, path string) (*Stats, error) {
	type nft struct {
		Stats Stats `json:"stats"`
	}
	type collection struct {
		NFT nft `json:"collection"`
	}

	result := collection{}
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
	return &result.NFT.Stats, nil
}
