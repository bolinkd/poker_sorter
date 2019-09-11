package models

type Cards []*Card

func (cs Cards) ToString() string {
	rtn := ""
	for _, card := range cs {
		rtn += card.toString() + ", "
	}
	return rtn[:len(rtn)-2] + "\n"
}
