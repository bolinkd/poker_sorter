package main

import (
	"fmt"
	"github.com/bolinkd/poker_sorter/models"
)

func main() {
	cards := models.Cards{}
	for suit := models.Heart; suit <= models.Spade; suit++ {
		for value := models.One; value <= models.King; value++ {
			cards = append(cards, &models.Card{
				Suit:  suit,
				Value: value,
			})
		}
	}

	fmt.Println(cards.ToString())
}
