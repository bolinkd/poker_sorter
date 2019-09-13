package main

import (
	"fmt"
	"github.com/bolinkd/poker_sorter/generation"
	"github.com/bolinkd/poker_sorter/models"
)

func comparePossibleHands(possibleHands1 models.Hands, possibleHands2 models.Hands) {
	for i, hand1 := range possibleHands1 {
		if i > len(possibleHands2) {
			fmt.Printf("Hand 1's %v Beats Hand 2's %v\n", hand1.ToString(), possibleHands2[len(possibleHands2)-1].ToString())
			return
		}

		hand2 := possibleHands2[i]
		cmp := hand1.Compare(hand2)
		if cmp > 0 {
			fmt.Printf("Hand 1's %v Beats Hand 2's %v\n", hand1.ToString(), hand2.ToString())
			return
		} else if cmp < 0 {
			fmt.Printf("Hand 2's %v Beats Hand 1's %v\n", hand2.ToString(), hand1.ToString())
			return
		} else {
			fmt.Printf("Users both have the same value cards: (%v and %v)\n", hand1.ToString(), hand2.ToString())
		}

	}

	if len(possibleHands1) == len(possibleHands2) {
		fmt.Println("Users have the exact same hand")
	}
}

func main() {
	deck := generation.GenerateDeck()

	hand1, err := generation.GetHand(&deck)
	if err != nil {
		return
	}
	// hand1 = models.Cards_MatchingPair1
	fmt.Println("Hand 1: " + hand1.ToString())
	possibleHands1 := hand1.GetPossibleHands()

	hand2, err := generation.GetHand(&deck)
	if err != nil {
		return
	}
	// hand2 = models.Cards_MatchingPair2
	fmt.Println("Hand 2: " + hand2.ToString())
	possibleHands2 := hand2.GetPossibleHands()

	fmt.Println("---------------------------------------------------------------------")

	comparePossibleHands(possibleHands1, possibleHands2)

	fmt.Println("---------------------------------------------------------------------")
}
