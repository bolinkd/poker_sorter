package models

type HandType int

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

const (
	UnknownHandType = iota
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

func (h *Hand) updateBestHand(newHandType HandType, hand Cards) {
	if newHandType > h.BestHandType {
		h.BestHandType = newHandType
		h.BestHand = hand
	}
}

func (h *Hand) Evaluate() error {
	// Check for cards with matching suits
	if ok, matchingSuits := h.Cards.hasMatchingSuit(); ok {
		h.updateBestHand(Flush, matchingSuits)
	}

	// Check for cards in sequence
	if ok, sequentialCards := h.Cards.hasSequence(); ok {
		isRoyal := false
		isFlush := false
		// Check if royal (ie: Last card is King)
		if h.Cards[len(h.Cards)-1].Value == King {
			isRoyal = true
		}

		// check if flush
		if h.BestHandType == Flush {
			isFlush = true
		}

		// assign accordingly
		if isRoyal && isFlush {
			h.updateBestHand(RoyalFlush, sequentialCards)
			return nil
		} else if isFlush {
			h.updateBestHand(StraightFlush, sequentialCards)
			return nil
		} else {
			h.updateBestHand(Straight, sequentialCards)
		}

	}

	// Check for Cards with matching Values
	if ok, matchingCards := h.Cards.hasMatchingValues(); ok {
		switch len(matchingCards) {
		case 2: // One or more sets of matching cards
			{
				c1 := matchingCards[0]
				c2 := matchingCards[1]

				if len(c1) == 2 && len(c2) == 2 {
					h.updateBestHand(TwoPair, append(c1, c2...))

				} else {
					// Must be full house
					h.updateBestHand(FullHouse, append(c1, c2...))
				}
				break
			}
		case 1: // Only one set of matching cards
			{
				cards := matchingCards[0]
				switch len(cards) {
				case 2:
					{
						h.updateBestHand(OnePair, cards)
					}
				case 3:
					{
						h.updateBestHand(ThreeOfAKind, cards)
					}
				case 4:
					{
						h.updateBestHand(FourOfAKind, cards)
					}
				}
			}
		}
		// Always Better Than High Card
		return nil
	}
	if ok, cards := h.Cards.hasHighCard(); ok {
		h.updateBestHand(HighCard, cards)
		return nil
	}
	return nil
}
