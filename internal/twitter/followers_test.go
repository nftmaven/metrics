package twitter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	fday  = "2022-07-16"
	fslug = "xoxonft"
)

func TestParseFollowersSuccess(t *testing.T) {

	expected := FollowerData{
		Day:        fday,
		Slug:       fslug,
		Followers:  33385,
		Following:  314,
		TweetCount: 810,
		UserName:   "xoxonft_io",
	}
	actual, err := ParseFollowers(fday, fslug, "test_data/xoxonft.json")
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, expected, *actual)
}
