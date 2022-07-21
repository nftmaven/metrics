package twitter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	day     = "2022-07-17"
	chain   = "ethereum"
	slug    = "regulars"
	dsource = "opensea"
)

func TestParseSuccess(t *testing.T) {

	expected := TwitterStats{
		Day:        day,
		Criterion:  chain,
		Slug:       slug,
		DataSource: dsource,
		SearchHits: 8,
		PublicMetrics: PublicMetrics{
			RetweetCount: 218,
			ReplyCount:   6,
			LikeCount:    29,
			QuoteCount:   5,
		},
	}
	actual, err := ParseSearchStats(chain, day, dsource, slug, "test_data/1")
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, expected, *actual)
}

func TestParseFailedToRead(t *testing.T) {

	actual, err := ParseSearchStats(chain, day, dsource, slug, "not_there")
	assert.Nil(t, actual)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "no such file or directory")
}
