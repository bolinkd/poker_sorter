package generation

import (
	"github.com/bolinkd/poker_sorter/models"
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	s := rand.NewSource(time.Now().UnixNano())
	r = rand.New(s)
}

func GenerateDeck() models.Cards {
	deck := models.Cards{}

	for suit := models.Heart; suit <= models.Spade; suit++ {
		for value := models.Two; value <= models.Ace; value++ {
			deck = append(deck, &models.Card{
				Suit:  suit,
				Value: value,
			})
		}
	}

	return deck
}

func Shuffle(cards *models.Cards) {
	r.Shuffle(len(*cards), func(i, j int) { (*cards)[i], (*cards)[j] = (*cards)[j], (*cards)[i] })
}

func GetHand(deck *models.Cards) (models.Cards, error) {
	Shuffle(deck)
	return (*deck)[:5], nil
}
