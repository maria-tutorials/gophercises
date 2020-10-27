package deck

import (
	"fmt"
	"testing"
)

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

func TestNew(t *testing.T) {
	cards := New()
	fmt.Println(cards)
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in the deck")
	}
}

func TestDefaultSort(t *testing.T) {
	d := New(DefaultSort)
	c := Card{Rank: Ace, Suit: Spade}

	if d[0] != c {
		t.Error("Expected Ace of Spades. Received", d[0])
	}
}
