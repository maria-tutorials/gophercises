//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
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

func Shuffle1(d Deck) Deck {
	now := time.Now().UnixNano()
	seed := rand.NewSource(now)
	r := rand.New(seed)

	for i := range d {
		newpos := r.Intn(len(d) - 1)

		d[i], d[newpos] = d[newpos], d[i] //weird-magic swap
	}

	return d
}

func Shuffle2(d Deck) Deck {
	s := make(Deck, len(d))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(d))

	for i, j := range perm {
		s[i] = d[j]
	}
	return s
}

//New(Jokers(3))
func Jokers(n int) func(Deck) Deck {
	return func(d Deck) Deck {
		for i := 0; i < n; i++ {
			d = append(d, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}
		return d
	}
}

func Filter(f func(c Card) bool) func(Deck) Deck {
	return func(d Deck) Deck {
		ret := Deck{}

		for _, c := range d {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

func NumberDecks(n int) func(Deck) Deck {
	return func(d Deck) Deck {
		ret := Deck{}
		for i := 0; i < n; i++ {
			ret = append(ret, d...)
		}
		return ret
	}
}
