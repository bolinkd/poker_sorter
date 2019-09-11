package models

type Suit int

const (
	UnknownSuit Suit = iota
	Heart
	Diamond
	Club
	Spade
)

func (s Suit) toString() string {
	switch s {
	case Heart:
		return "♥"
	case Diamond:
		return "♦"
	case Club:
		return "♣"
	case Spade:
		return "♠"
	default:
		return "U"
	}
}
