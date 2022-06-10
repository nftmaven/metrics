package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testPath       = "test_data/060_2022-06-10_node-monkey-genesis.json"
	brokenJSONPath = "test_data/broken.json"
)

func TestParseSuccess(t *testing.T) {

	nft, err := Parse("opensea", "solana", testPath)
	assert.Nil(t, err)
	assert.NotNil(t, nft)
	assert.Equal(t, "opensea", nft.DataSource)
	assert.Equal(t, "solana", nft.Chain)
	assert.Equal(t, "node-monkey-genesis", nft.Slug)
	assert.Equal(t, "Node Monkey Genesis", nft.Name)
	assert.Equal(t, "https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2FaQuHZ7yM7lT1ZXkroTqIuKYmmR8XbMP7t5JYtu1uzjdaLzImxbhNJXpkB_Zb1JtyXuUeNovOH_7mVpbDYPvQHUNgtVkFQeiKCzUSdA%3Ds10000?fit=max&h=300&w=300&auto=format&s=0f4def3ab568e4538d43827a1b9e85b7", nft.ImageURL)
	assert.Equal(t, "https://discord.gg/nodemonkey", nft.DiscordURL)
	assert.Equal(t, "", nft.URL)
	assert.Equal(t, "approved", nft.SafelistStatus)
}

func TestParseFailedToRead(t *testing.T) {

	nft, err := Parse("opensea", "solana", "/a/b/c")
	assert.Nil(t, nft)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "failed to read ")
}

func TestParseFailedToParse(t *testing.T) {

	nft, err := Parse("opensea", "solana", brokenJSONPath)
	assert.Nil(t, nft)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "failed to parse ")
}
