package models

type Card struct {
	Suit  Suit  `json:"suit"`
	Value Value `json:"value"`
}

func (c *Card) toString() string {
	return c.Value.toString() + c.Suit.toString()
}

func (c *Card) Compare(c2 *Card) int {
	return int(c.Value - c2.Value)
}
