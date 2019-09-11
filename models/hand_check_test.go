package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasMatchingValues_FourOfAKindValid(t *testing.T) {
	assert := assert.New(t)
	ok, matchingCards := cards_FourOfAKind.hasMatchingValues()
	assert.True(ok, "Should return that values match")
	assert.Equal(len(matchingCards), 1, "Should return a match for a single value")
	assert.Equal(len(matchingCards[0]), 4, "Should return 4 matching cards for the value")
}

func Test_hasMatchingValues_TwoPairValid(t *testing.T) {
	assert := assert.New(t)
	ok, matchingCards := cards_TwoPair.hasMatchingValues()
	assert.True(ok, "Should return that values match")
	assert.Equal(len(matchingCards), 2, "Should return a match for two different values")
	assert.Equal(len(matchingCards[0]), 2, "Should return 2 matching cards for the first value")
	assert.Equal(len(matchingCards[1]), 2, "Should return 2 matching cards for the second value")
}

func Test_hasMatchingValues_Invalid(t *testing.T) {
	assert := assert.New(t)
	ok, _ := cards_NoPairs.hasMatchingValues()
	assert.False(ok, "Should return that no values match")
}

func Test_hasHighCard_Valid(t *testing.T) {
	assert := assert.New(t)
	ok, highCard := cards_ValidHighCard.hasHighCard()
	assert.True(ok, "Should return a valid high card")
	assert.Equal(highCard.Value, King, "Should return the King")
}

func Test_hasHighCard_Invalid(t *testing.T) {
	assert := assert.New(t)
	ok, _ := cards_InvalidHighCard.hasHighCard()
	assert.False(ok, "Should return no valid cards")
}
