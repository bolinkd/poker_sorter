package models

import (
	"fmt"
	"sort"
)

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
	RelevantCards Cards    `json:"best_cards"`
	BestHandType  HandType `json:"hand_type"`
}

type Hands []*Hand

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

func (h *Hand) ToString() string {
	return h.BestHandType.ToString() + ": (" + h.RelevantCards.ToString() + ")"
}

func (h *Hand) compareFullHouse(h2 *Hand) int {
	hand1Map := make(map[Value]Cards)
	hand2Map := make(map[Value]Cards)

	for _, card := range h.RelevantCards {
		hand1Map[card.Value] = append(hand1Map[card.Value], card)
	}

	for _, card := range h2.RelevantCards {
		hand2Map[card.Value] = append(hand2Map[card.Value], card)
	}

	var hand1Triple, hand2Triple Cards
	var hand1Pair, hand2Pair Cards
	for _, cards := range hand1Map {
		if len(cards) == 3 {
			hand1Triple = cards
		} else {
			hand1Pair = cards
		}
	}

	for _, cards := range hand2Map {
		if len(cards) == 3 {
			hand2Triple = cards
		} else {
			hand2Pair = cards
		}
	}

	cmp := 0
	if hand1Triple != nil && hand2Triple != nil {
		cmp = hand1Triple[0].Compare(hand2Triple[0])
		if cmp == 0 {
			if hand1Pair != nil && hand2Pair != nil {
				fmt.Println(hand1Pair.ToString(), hand2Pair.ToString())
				cmp = hand1Pair[0].Compare(hand2Pair[0])
			}
		}
	}

	return cmp
}

func (h *Hand) Compare(h2 *Hand) int {
	if h.BestHandType != h2.BestHandType {
		return int(h.BestHandType - h2.BestHandType)
	}

	if h.BestHandType == FullHouse {
		return h.compareFullHouse(h2)
	}

	sort.SliceStable(h.RelevantCards, func(i int, j int) bool {
		return h.RelevantCards[i].Value > h.RelevantCards[j].Value
	})
	sort.SliceStable(h2.RelevantCards, func(i int, j int) bool {
		return h2.RelevantCards[i].Value > h2.RelevantCards[j].Value
	})

	return h.RelevantCards[0].Compare(h2.RelevantCards[0])
}

func (hs Hands) ToString() string {
	rtn := ""
	if len(hs) == 0 {
		return rtn
	}
	for _, card := range hs {
		rtn += card.ToString() + "\n"
	}
	return rtn[:len(rtn)-2]
}
