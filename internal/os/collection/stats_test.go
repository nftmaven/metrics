package collection

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestParseStatsSuccess(t *testing.T) {

	stats, err := ParseStats("opensea", "solana", testPath)
	assert.Nil(t, err)
	assert.NotNil(t, stats)
	assert.Equal(t, decimal.RequireFromString("6.7207040888"), stats.Day1Volume)
	assert.Equal(t, decimal.RequireFromString("1.2364925373134326"), stats.Day1Change)
	assert.Equal(t, decimal.RequireFromString("6.0"), stats.Day1Sales)
	assert.Equal(t, decimal.RequireFromString("1.1201173481333333"), stats.Day1AvgPrice)
	assert.Equal(t, decimal.RequireFromString("17.372826088799997"), stats.Day7Volume)
	assert.Equal(t, decimal.RequireFromString("0.6875939440148126"), stats.Day7Change)
	assert.Equal(t, decimal.RequireFromString("17.0"), stats.Day7Sales)
	assert.Equal(t, decimal.RequireFromString("1.1201173481333333"), stats.Day1AvgPrice)
	assert.Equal(t, decimal.RequireFromString("70.23271693872721"), stats.Day30Volume)
	assert.Equal(t, decimal.RequireFromString("-0.15670126052170605"), stats.Day30Change)
	assert.Equal(t, decimal.RequireFromString("63.0"), stats.Day30Sales)
	assert.Equal(t, decimal.RequireFromString("1.1148050307734478"), stats.Day30AvgPrice)
	assert.Equal(t, decimal.RequireFromString("371.1233217792617"), stats.TotalVolume)
	assert.Equal(t, decimal.RequireFromString("466.0"), stats.TotalSales)
	assert.Equal(t, decimal.RequireFromString("250.0"), stats.Count)
	assert.Equal(t, decimal.RequireFromString("190"), stats.NumOwners)
	assert.Equal(t, decimal.RequireFromString("0.7964019780670851"), stats.AvgPrice)
	assert.Equal(t, decimal.RequireFromString("255.48273659999995"), stats.MarketCap)
	assert.Equal(t, decimal.Decimal{}, stats.FloorPrice)
}

func TestParseStatsFailedToRead(t *testing.T) {

	nft, err := Parse("opensea", "solana", "/a/b/c")
	assert.Nil(t, nft)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "failed to read ")
}

func TestParseStatsFailedToParse(t *testing.T) {

	nft, err := Parse("opensea", "solana", brokenJSONPath)
	assert.Nil(t, nft)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "failed to parse ")
}
