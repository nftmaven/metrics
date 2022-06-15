package collection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/shopspring/decimal"
)

type Stats struct {
	Day1Volume    decimal.Decimal `json:"one_day_volume" db:"d1_volume"`
	Day1Change    decimal.Decimal `json:"one_day_change" db:"d1_change"`
	Day1Sales     decimal.Decimal `json:"one_day_sales" db:"d1_sales"`
	Day1AvgPrice  decimal.Decimal `json:"one_day_average_price" db:"d1_avg_price"`
	Day7Volume    decimal.Decimal `json:"seven_day_volume" db:"d7_volume"`
	Day7Change    decimal.Decimal `json:"seven_day_change" db:"d7_change"`
	Day7Sales     decimal.Decimal `json:"seven_day_sales" db:"d7_sales"`
	Day7AvgPrice  decimal.Decimal `json:"seven_day_average_price" db:"d7_avg_price"`
	Day30Volume   decimal.Decimal `json:"thirty_day_volume" db:"d30_volume"`
	Day30Change   decimal.Decimal `json:"thirty_day_change" db:"d30_change"`
	Day30Sales    decimal.Decimal `json:"thirty_day_sales" db:"d30_sales"`
	Day30AvgPrice decimal.Decimal `json:"thirty_day_average_price" db:"d30_avg_price"`
	TotalVolume   decimal.Decimal `json:"total_volume" db:"total_volume"`
	TotalSales    decimal.Decimal `json:"total_sales" db:"total_sales"`
	TotalSupply   decimal.Decimal `json:"total_supply" db:"total_supply"`
	Owners        decimal.Decimal `json:"num_owners" db:"owners"`
	AvgPrice      decimal.Decimal `json:"average_price" db:"avg_price"`
	MarketCap     decimal.Decimal `json:"market_cap" db:"market_cap"`
	FloorPrice    decimal.Decimal `json:"floor_price" db:"floor_price"`
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
