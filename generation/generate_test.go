package generation

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

const DECK_LENGTH = 52

func TestGeneratorCreatesDeckWith52Cards(t *testing.T) {
	deck := GenerateDeck()
	assert.Equal(t, len(deck), DECK_LENGTH, "Deck should contain 52 cards")
}

func TestHandGeneration(t *testing.T) {
	assert := assert.New(t)

	deck := GenerateDeck()

	r = rand.New(rand.NewSource(99))

	hand, err := GetHand(&deck)

	assert.Equal(len(hand), 5, "Hand should contain 5 cards")
	assert.NoError(err, "Should return a successful hand draw")
}
