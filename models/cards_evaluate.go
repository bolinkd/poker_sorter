package models

import "fmt"

func (cs *Cards) GetPossibleHands() Hands {
	// Check for cards with matching suits
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
		fmt.Println(possibleHands.ToString())
	}
	/*
		// Check what high card is
		if ok, highCard := h.Cards.hasHighCard(); ok {
			h.updateBestHand(HighCard, Cards{highCard})
			return
		}
	*/
	return nil
}
