package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasMatchingValues_FourOfAKindValid(t *testing.T) {
	assert := assert.New(t)
	ok, matchingCards := Cards_FourOfAKind.hasMatchingValues()
	assert.True(ok, "Should return that values match")
	assert.Equal(len(matchingCards), 1, "Should return a match for a single value")
	assert.Equal(len(matchingCards[0]), 4, "Should return 4 matching cards for the value")
}

func Test_hasMatchingValues_TwoPairValid(t *testing.T) {
	assert := assert.New(t)
	ok, matchingCards := Cards_TwoPair.hasMatchingValues()
	assert.True(ok, "Should return that values match")
	assert.Equal(len(matchingCards), 2, "Should return a match for two different values")
	assert.Equal(len(matchingCards[0]), 2, "Should return 2 matching cards for the first value")
	assert.Equal(len(matchingCards[1]), 2, "Should return 2 matching cards for the second value")
}

func Test_hasMatchingValues_Invalid(t *testing.T) {
	assert := assert.New(t)
	ok, _ := Cards_NoPairs.hasMatchingValues()
	assert.False(ok, "Should return that no values match")
}
