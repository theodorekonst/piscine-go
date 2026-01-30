package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	cardsPerPlayer := 3 // πάντα 3 (12 κάρτες / 4 παίκτες)

	for player := 0; player < 4; player++ {
		start := player * cardsPerPlayer
		end := start + cardsPerPlayer

		fmt.Printf("Player %d: ", player+1)

		for i := start; i < end; i++ {
			if i < end-1 {
				fmt.Printf("%d, ", deck[i])
			} else {
				fmt.Printf("%d", deck[i])
			}
		}
		fmt.Printf("\n")
	}
}
