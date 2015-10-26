package poker

import (
	"fmt"
	"strconv"
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

type Suit string

const (
	Clubs    = Suit("Clubs")
	Diamonds = Suit("Diamonds")
	Hearts   = Suit("Hearts")
	Spades   = Suit("Spades")
)

type Card struct {
	Rank Rank
	Suit Suit
}

func (card Card) Code() string {
	return fmt.Sprintf("%s%c", card.Rank, card.Suit[0])
}
