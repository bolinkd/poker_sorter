package models

import (
	"fmt"
	"sort"
)

func (cs *Cards) GetPossibleHands() Hands {
	// Sort by value
	sort.SliceStable(*cs, func(i int, j int) bool {
		return (*cs)[i].Value > (*cs)[j].Value
	})

	// Check for cards with matching suits
	allHands := make(Hands, 0)
	// TODO: Sequences and Flushes
	/*
		if ok, matchingSuits := cs.getMatchingSuits(); ok {

		}

		// Check for cards in sequence
		if ok, sequentialCards := cs.getSequences(); ok {
			isRoyal := false
			isFlush := false
			// Check if royal (ie: Last card is King)
			if h.Cards[len(h.Cards)-1].Value == Ace {
				isRoyal = true
			}

			// check if flush
			if h.BestHandType == Flush {
				isFlush = true
			}

			if isRoyal && isFlush {
				h.updateBestHand(RoyalFlush, sequentialCards)
				return
			} else if isFlush {
				h.updateBestHand(StraightFlush, sequentialCards)
				return
			} else {
				h.updateBestHand(Straight, sequentialCards)
			}

		}
	*/

	// Check for Cards with matching Values
	if possibleHands, ok := cs.groupMatchingValues(); ok {
		allHands = append(allHands, possibleHands...)
	}

	if possibleHands, ok := cs.groupHighCards(); ok {
		allHands = append(allHands, possibleHands...)
	}

	sort.SliceStable(allHands, func(i int, j int) bool {
		return allHands[i].BestHandType > allHands[j].BestHandType
	})

	fmt.Println(allHands.ToString())
	return allHands
}
