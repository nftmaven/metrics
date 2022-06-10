package top100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testPath      = "test_data/ethereum.csv"
	brokenCSVPath = "test_data/broken.csv"
)

func TestProcessSuccess(t *testing.T) {

	data, err := Process("opensea", "ethereum", testPath)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, 100, len(data))
	sr := data[99]
	assert.Equal(t, 100, sr.Rank)
	assert.Equal(t, "2022-06-10", sr.Day)
	assert.Equal(t, "Dippies Vans", sr.Name)
	assert.Equal(t, "vansofficial", sr.Slug)
	assert.Equal(t, false, sr.IsVerified)
	sr = data[0]
	assert.Equal(t, 1, sr.Rank)
	assert.Equal(t, "2022-06-10", sr.Day)
	assert.Equal(t, "Otherdeed for Otherside", sr.Name)
	assert.Equal(t, "otherdeed", sr.Slug)
	assert.Equal(t, true, sr.IsVerified)
}

func TestProcessFailedToOpen(t *testing.T) {

	nft, err := Process("opensea", "solana", "/a/b/c")
	assert.Nil(t, nft)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "failed to open ")
}

func TestProcessFailedToProcess(t *testing.T) {

	nft, err := Process("opensea", "solana", brokenCSVPath)
	assert.Nil(t, nft)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "failed to unmarshal ")
}
