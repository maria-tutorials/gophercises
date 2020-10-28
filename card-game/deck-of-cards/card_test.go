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

func TestJokers(t *testing.T) {
	numJokers := 3

	deck := New(Jokers(numJokers))
	n := 0

	for _, c := range deck {
		if c.Suit == Joker {
			n++
		}
	}

	if n != numJokers {
		t.Error("Expected", numJokers, "Jokers. Received", n)
	}
}

func TestFiler(t *testing.T) {
	filter := func(c Card) bool {
		return c.Rank == Two || c.Rank == Three
	}

	d := New(Filter(filter))

	for _, c := range d {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Found some Twos and Threes")
		}
	}

}

func TestDeck(t *testing.T) {
	d := New(NumberDecks(3))

	if len(d) != 13*4*3 {
		t.Errorf("Expected %d cards, received %d", 13*4*3, len(d))
	}
}
