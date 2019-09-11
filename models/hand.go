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
			if h1.BestHand[0].Value == h2.BestHand[0].Value {
				return false, ErrHandsAreEquivalent
			}
			return h1.BestHand[0].Value > h2.BestHand[0].Value, nil
		}
		return false, ErrHandsAreEquivalent
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
