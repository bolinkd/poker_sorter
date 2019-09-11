package main

import (
	"fmt"
	"github.com/bolinkd/poker_sorter/generation"
	"github.com/bolinkd/poker_sorter/models"
)

func main() {
	deck := generation.GenerateDeck()
	// fmt.Println(deck.ToString())

	hand, err := generation.GetHand(&deck)
	if err != nil {
		// TODO: ERROR
	}
	fmt.Println(hand.Cards.ToString())
	err = hand.Evaluate()
	fmt.Println(hand.BestHand.ToString(), hand.BestHandType)

	hand2 := models.NewHand(models.Cards{
		&models.Card{
			1, 1,
		},
		&models.Card{
			1, 2,
		},
		&models.Card{
			1, 3,
		},
		&models.Card{
			1, 4,
		},
		&models.Card{
			1, 5,
		},
	})
	if err != nil {
		// TODO: ERROR
	}
	fmt.Println(hand2.Cards.ToString())
	err = hand2.Evaluate()
	fmt.Println(hand2.BestHand.ToString(), hand2.BestHandType)

}
