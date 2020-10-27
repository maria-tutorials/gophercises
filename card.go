//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"sort"
)

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

const (
	minRank = Ace
	maxRank = King
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

type Deck []Card

func New(opts ...func(Deck) Deck) Deck {
	d := Deck{}

	for suit := 0; suit < 4; suit++ {
		for rank := minRank; rank <= maxRank; rank++ {
			d = append(d, Card{Suit: Suit(suit), Rank: rank})
		}
	}

	for _, opt := range opts {
		d = opt(d)
	}

	return d
}

func absCardRank(c Card) int {
	return int(c.Suit) * int(maxRank) * int(c.Rank)
}

func DefaultSort(d Deck) Deck {
	//sort.Sort(ByRank(d))
	sort.Slice(d, ByRank(d).Less)
	return d
}

// ByRank using sort magic
type ByRank Deck

func (a ByRank) Len() int           { return len(a) }
func (a ByRank) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool { return absCardRank(a[i]) < absCardRank(a[j]) }
