//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

const (
	Red   string = "Red"
	Black string = "Black"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit
	Color string
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String() //Jokers have no extras
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}
