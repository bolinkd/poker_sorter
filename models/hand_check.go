package models

import (
	"sort"
)

func (cs Cards) hasMatchingValues() (bool, []Cards) {
	cardMap := make(map[Value]Cards)
	for _, card := range cs {
		if cardMap[card.Value] == nil {
			cardMap[card.Value] = Cards{card}
		} else {
			cardMap[card.Value] = append(cardMap[card.Value], card)
		}
	}

	var matchingValueCards []Cards
	for _, cards := range cardMap {
		if len(cards) >= 2 {
			matchingValueCards = append(matchingValueCards, cards)
		}
	}

	return matchingValueCards != nil, matchingValueCards
}

func (cs Cards) hasHighCard() (bool, *Card) {
	if len(cs) == 0 {
		return false, nil
	}
	var maxCard *Card
	for _, card := range cs {
		if maxCard == nil || card.Value > maxCard.Value {
			maxCard = card
		}
	}
	return true, maxCard
}

func (cs Cards) hasMatchingSuit() (bool, Cards) {
	suit := cs[0].Suit
	for _, card := range cs[1:] {
		if card.Suit != suit {
			return false, nil
		}
	}
	return true, cs
}

func (cs Cards) hasSequence() (bool, Cards) {
	sort.SliceStable(cs, func(i, j int) bool {
		return cs[i].Value < cs[j].Value
	})

	seqVal := cs[0].Value
	for _, card := range cs[1:] {
		if card.Value == seqVal+1 {
			seqVal++
		} else {
			return false, Cards{}
		}
	}
	return true, cs
}
