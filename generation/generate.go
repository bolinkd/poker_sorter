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
		for value := models.One; value <= models.King; value++ {
			deck = append(deck, &models.Card{
				Suit:  suit,
				Value: value,
			})
		}
	}

	return deck
}

func ShuffleDeck() {

}

func valueInSlice(val int, list []int) bool {
	for _, item := range list {
		if item == val {
			return true
		}
	}
	return false
}

func getRandomNumber(min int, max int) int {
	return r.Intn(max-min+1) + min
}

func getRandomNumbers(count int, min int, max int) []int {
	indexes := []int{}
	for len(indexes) < count {
		newIndex := getRandomNumber(min, max)
		if !valueInSlice(newIndex, indexes) {
			indexes = append(indexes, newIndex)
		}
	}
	// List numbers in descending order
	return indexes
}

func GetHand(deck *models.Cards) (models.Hand, error) {
	indexes := getRandomNumbers(5, 0, len(*deck)-1)
	cards := models.Cards{}
	for _, idx := range indexes {
		cards = append(cards, (*deck)[idx])
	}
	return models.NewHand(cards), nil
}
