package models

import (
	"errors"
)

var ErrHandsAreEquivalent = errors.New("hands are equivalent")

type HandType int

const (
	UnknownHandType HandType = iota
	HighCard
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

type Hand struct {
	Cards        Cards    `json:"cards"`
	BestHand     Cards    `json:"best_cards"`
	BestHandType HandType `json:"hand_type"`
}

func NewHand(cards Cards) Hand {
	return Hand{
		Cards:        cards,
		BestHand:     make(Cards, 0),
		BestHandType: UnknownHandType,
	}
}

func compareMaxValues(c1 Cards, c2 Cards) (bool, error) {
	maxC1 := c1[0].Value
	for _, card := range c1[1:] {
		if card.Value > maxC1 {
			maxC1 = card.Value
		}
	}

	maxC2 := c2[0].Value
	for _, card := range c2[1:] {
		if card.Value > maxC2 {
			maxC2 = card.Value
		}
	}
	if maxC1 > maxC2 {
		return true, nil
	} else if maxC2 > maxC1 {
		return false, nil
	} else {
		return false, ErrHandsAreEquivalent
	}
}

func compareMatchingValues(c1 Cards, c2 Cards) (bool, error) {
	if c1[0].Value == c2[0].Value {
		return false, ErrHandsAreEquivalent
	}
	return c1[0].Value > c2[0].Value, nil
}

func (h1 Hand) IsBetterThan(h2 Hand) (bool, error) {
	if h1.BestHandType > h2.BestHandType {
		return true, nil
	} else if h1.BestHandType < h2.BestHandType {
		return false, nil
	} else {
		// Equvalent Type
		switch h1.BestHandType {
		case HighCard:
			fallthrough
		case OnePair:
			fallthrough
		case ThreeOfAKind:
			fallthrough
		case FourOfAKind:
			return compareMatchingValues(h1.BestHand, h2.BestHand)
		case TwoPair:
			fallthrough
		case Straight:
			fallthrough
		case StraightFlush:
			fallthrough
		case Flush:
			fallthrough
		case FullHouse:
			return compareMaxValues(h1.BestHand, h2.BestHand)
			// Royal Flush (has to be the same)
		default:
			return false, ErrHandsAreEquivalent
		}
	}
}

func (h HandType) ToString() string {
	switch h {
	case HighCard:
		return "High Card"
	case OnePair:
		return "One Pair"
	case TwoPair:
		return "Two Pair"
	case ThreeOfAKind:
		return "Three Of A Kind"
	case Straight:
		return "Straight"
	case Flush:
		return "Flush"
	case FullHouse:
		return "Full House"
	case FourOfAKind:
		return "Four Of A Kind"
	case StraightFlush:
		return "Straight Flush"
	case RoyalFlush:
		return "Royal Flush"
	default:
		return "Unknown"
	}
}
