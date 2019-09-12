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

func (cs Cards) groupHighCards() (Hands, bool) {
	possibleHands := make(Hands, 0)
	for _, card := range cs {
		possibleHands = append(possibleHands, &Hand{
			BestHandType:  HighCard,
			RelevantCards: Cards{card},
		})
	}
	return possibleHands, len(possibleHands) != 0
}

func (cs Cards) groupSequences() (Hands, bool) {
	// TODO: Expand for more than 5 cards
	// TODO: Check for Ace as Low
	prevCard := cs[0]
	hands := make(Hands, 0)
	for _, card := range cs[1:] {
		if card.Value == prevCard.Value-1 {
			prevCard = card
		} else {
			return hands, false
		}
	}
	hand := &Hand{
		BestHandType:  Straight,
		RelevantCards: cs,
	}
	hands = append(hands, hand)

	// Check straight flush
	isFlush := true
	suit := hand.RelevantCards[0].Suit
	for _, card := range hands[0].RelevantCards[1:] {
		if card.Suit != suit {
			isFlush = false
			break
		}
	}

	// Check royal flush
	isRoyal := hand.RelevantCards[0].Value == Ace
	if isFlush {
		hands = append(hands, &Hand{
			RelevantCards: hand.RelevantCards,
			BestHandType:  StraightFlush,
		})
	}
	if isFlush && isRoyal {
		hands = append(hands, &Hand{
			RelevantCards: hand.RelevantCards,
			BestHandType:  RoyalFlush,
		})
	}
	return hands, true
}

func (cs Cards) groupMatchingSuits() (Hands, bool) {
	// TODO: Expand for more than 5 cards
	suitMap := make(map[Suit]Cards)
	hands := make(Hands, 0)
	for _, card := range cs {
		if suitMap[card.Suit] == nil {
			suitMap[card.Suit] = Cards{card}
		} else {
			suitMap[card.Suit] = append(suitMap[card.Suit], card)
		}
	}

	for _, suitCards := range suitMap {
		if len(suitCards) == 5 {
			hands = append(hands, &Hand{
				BestHandType:  Flush,
				RelevantCards: suitCards,
			})
		}
	}

	return hands, len(hands) != 0
}
