package main

import (
	"fmt"
	"github.com/bolinkd/poker_sorter/generation"
	"github.com/bolinkd/poker_sorter/models"
)

func CompareHands(hand1 models.Hand, hand2 models.Hand) {
	better, err := hand1.IsBetterThan(hand2)
	if err != nil {
		fmt.Printf("Hand 1's %v is the same as Hand 2's %v\n", hand1.BestHandType.ToString(), hand2.BestHandType.ToString())
		return
	}

	if better {
		fmt.Printf("Hand 1's %v beats Hand 2's %v\n", hand1.BestHandType.ToString(), hand2.BestHandType.ToString())
	} else {
		fmt.Printf("Hand 2's %v beats Hand 1's %v\n", hand2.BestHandType.ToString(), hand1.BestHandType.ToString())
	}
}

func main() {
	deck := generation.GenerateDeck()

	hand1, err := generation.GetHand(&deck)
	if err != nil {
		return
	}
	fmt.Println("Hand 1: " + hand1.Cards.ToString())
	hand1.Evaluate()

	hand2, err := generation.GetHand(&deck)
	if err != nil {
		return
	}
	fmt.Println("Hand 2: " + hand2.Cards.ToString())
	hand2.Evaluate()

	CompareHands(hand1, hand2)

	fmt.Println("Hand 1 Relevent Cards: ", hand1.BestHand.ToString())
	fmt.Println("Hand 2 Relevent Cards: ", hand2.BestHand.ToString())

}
