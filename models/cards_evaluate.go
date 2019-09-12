package models

import (
	"sort"
)

func (cs *Cards) GetPossibleHands() Hands {
	// Sort by value
	sort.SliceStable(*cs, func(i int, j int) bool {
		return (*cs)[i].Value > (*cs)[j].Value
	})

	// Check for cards with matching suits
	allHands := make(Hands, 0)

	// Flushes
	if flushHands, ok := cs.groupMatchingSuits(); ok {
		allHands = append(allHands, flushHands...)
	}

	// Sequences
	if sequenceHands, ok := cs.groupSequences(); ok {
		allHands = append(allHands, sequenceHands...)
	}

	// Check for Cards with matching Values
	if possibleHands, ok := cs.groupMatchingValues(); ok {
		allHands = append(allHands, possibleHands...)
	}

	if possibleHands, ok := cs.groupHighCards(); ok {
		allHands = append(allHands, possibleHands...)
	}

	sort.SliceStable(allHands, func(i int, j int) bool {
		if allHands[i].BestHandType == allHands[j].BestHandType {
			return allHands[i].RelevantCards[0].Value > allHands[j].RelevantCards[0].Value
		}
		return allHands[i].BestHandType > allHands[j].BestHandType
	})
	return allHands
}
