package poker

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Rank int

func (rank Rank) Strength() int {
	if rank == 1 {
		return 13
	}
	return int(rank) - 1
}

func (rank Rank) String() string {
	switch rank {
	case 1:
		return "A"
	case 10:
		return "T"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return strconv.Itoa(int(rank))
	}
}

type Suit int

const (
	Clubs = Suit(iota)
	Diamonds
	Hearts
	Spades
)

func (suit Suit) String() string {
	switch suit {
	case Clubs:
		return "clubs"
	case Diamonds:
		return "diamonds"
	case Hearts:
		return "hearts"
	case Spades:
		return "Spades"
	}
	return ""
}

type Card struct {
	Rank Rank
	Suit Suit
}

func (card Card) String() string {
	return fmt.Sprintf("%s%c", card.Rank, strings.ToUpper(card.Suit.String())[0])
}

var InvalidCard = errors.New("Invalid card")

func ParseCard(s string, v *Card) error {
	var rank Rank
	switch s[0] {
	case 'A':
		rank = 1
	case 'T':
		rank = 10
	case 'J':
		rank = 11
	case 'Q':
		rank = 12
	case 'K':
		rank = 13
	default:
		n, err := strconv.Atoi(s[:1])
		if err != nil {
			return InvalidCard
		}
		rank = Rank(n)
		if rank < 2 || 13 < rank {
			return InvalidCard
		}
	}
	var suit Suit
	switch s[1] {
	case 'C':
		suit = Clubs
	case 'D':
		suit = Diamonds
	case 'H':
		suit = Hearts
	case 'S':
		suit = Spades
	default:
		return InvalidCard
	}
	*v = Card{
		Rank: rank,
		Suit: suit,
	}
	return nil
}

type cardSorter struct {
	cards []Card
}

func (cs cardSorter) Len() int {
	return len(cs.cards)
}

func (cs cardSorter) Swap(i, j int) {
	cs.cards[i], cs.cards[j] = cs.cards[j], cs.cards[i]
}

func (cs cardSorter) Less(i, j int) bool {
	cards := cs.cards
	m, n := cards[i].Rank.Strength(), cards[j].Rank.Strength()
	// if m == n {
	// 	return cards[i].Suit > cards[j].Suit
	// }
	return m > n
}

func SortCards(cards []Card) {
	sorter := cardSorter{cards}
	sort.Sort(sorter)
}
