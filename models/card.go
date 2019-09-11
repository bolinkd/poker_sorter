package models

type Card struct {
	Suit  Suit  `json:"suit"`
	Value Value `json:"value"`
}

type Cards []*Card

func (c *Card) toString() string {
	return c.Value.toString() + c.Suit.toString()
}

func (cs Cards) ToString() string {
	rtn := ""
	for _, card := range cs {
		rtn += card.toString() + ", "
	}
	return rtn[:len(rtn)-2] + "\n"
}
