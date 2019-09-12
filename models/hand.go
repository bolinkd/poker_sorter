package models

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

func (h *Hand) toString() string {
	return h.BestHandType.ToString() + ": " + h.RelevantCards.ToString()
}

func (hs Hands) ToString() string {
	rtn := ""
	if len(hs) == 0 {
		return rtn
	}
	for _, card := range hs {
		rtn += card.toString() + ", "
	}
	return rtn[:len(rtn)-2]
}
