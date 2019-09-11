package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHand_Evaluate_FourOfAKind(t *testing.T) {
	assert := assert.New(t)
	hand := Hand(hand_FourOfAKind)
	hand.Evaluate()
	assert.Equal(hand.BestHandType, FourOfAKind, "Should return four of a kind")
	assert.Equal(len(hand.BestHand), 4, "Should return 4 cards")
}

func TestHand_Evaluate_TwoPair(t *testing.T) {
	assert := assert.New(t)
	hand := Hand(hand_TwoPair)
	hand.Evaluate()
	assert.Equal(hand.BestHandType, TwoPair, "Should return two pair")
	assert.Equal(len(hand.BestHand), 4, "Should return 4 cards")
}

func TestHand_Evaluate_FullHouse(t *testing.T) {
	assert := assert.New(t)
	hand := Hand(hand_FullHouse)
	hand.Evaluate()
	assert.Equal(hand.BestHandType, FullHouse, "Should return a full house")
	assert.Equal(len(hand.BestHand), 5, "Should return 5 cards")
}

func TestHand_Evaluate_RoyalFlush(t *testing.T) {
	assert := assert.New(t)
	hand := Hand(hand_RoyalFlush)
	hand.Evaluate()
	assert.Equal(hand.BestHandType, RoyalFlush, "Should return a royal flush")
	assert.Equal(len(hand.BestHand), 5, "Should return 5 cards")
}
