package main

import (
	"fmt"
	"github.com/bolinkd/poker_sorter/generation"
)

func main() {
	deck := generation.GenerateDeck()
	fmt.Println(deck.ToString())

	hand, err := generation.GetHand(&deck)
	if err != nil {
		// TODO: ERROR
	}
	fmt.Println(hand.ToString())

}
