package models

func (cs Cards) groupMatchingValues() (Hands, bool) {
	cardMap := make(map[Value]Cards)
	for _, card := range cs {
		if cardMap[card.Value] == nil {
			cardMap[card.Value] = Cards{card}
		} else {
			cardMap[card.Value] = append(cardMap[card.Value], card)
		}
	}

	possibleHands := make(Hands, 0)
	for _, matchingCards := range cardMap {
		hand := Hand{RelevantCards: matchingCards}
		switch len(matchingCards) {
		case 2:
			hand.BestHandType = OnePair
		case 3:
			hand.BestHandType = ThreeOfAKind
		case 4:
			hand.BestHandType = FourOfAKind
		default:
			break
		}
		if hand.BestHandType != UnknownHandType {
			possibleHands = append(possibleHands, &hand)
		}
	}

	// Check for two pairs and full house
	if len(possibleHands) >= 2 {
		for i, hand := range possibleHands[:len(possibleHands)-1] {
			prevHand := possibleHands[i+1]
			newHand := Hand{RelevantCards: append(hand.RelevantCards, prevHand.RelevantCards...)}
			if hand.BestHandType == OnePair && prevHand.BestHandType == OnePair {
				newHand.BestHandType = TwoPair
			} else if (hand.BestHandType == ThreeOfAKind && prevHand.BestHandType == OnePair) || (hand.BestHandType == OnePair && prevHand.BestHandType == ThreeOfAKind) {
				newHand.BestHandType = FullHouse
			}
			possibleHands = append(possibleHands, &newHand)
		}
	}

	return possibleHands, len(possibleHands) != 0
}

/*
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
	cardList := cs[1:]

	// If there is a two and an Ace it can count as a straight
	// TODO: Currently not working
	/*
		if seqVal == Two && cs[len(cs)-1].Value == Ace {
			cardList = cs[1 : len(cs)-1]
			cs = append(Cards{cs[len(cs)-1]}, cs[1:]...)
		}
*/
/*
	for _, card := range cardList {
		if card.Value == seqVal+1 {
			seqVal++
		} else {
			return false, Cards{}
		}
	}
	return true, cs
}

*/
