package deck

import "fmt"

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Diamond})
	fmt.Println(Card{Rank: Two, Suit: Heart})
	fmt.Println(Card{Rank: Nine, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	//Output:
	// Ace of Diamonds
	// Two of Hearts
	// Nine of Spades
	// Joker
}
