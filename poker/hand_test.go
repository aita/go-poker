package poker

import (
	"strings"
	"testing"
)

func parseCards(s string) []Card {
	cards := []Card{}
	for _, c := range strings.Split(s, " ") {
		var card Card
		if err := ParseCard(c, &card); err != nil {
			panic(err)
		}
		cards = append(cards, card)
	}
	return cards
}

func TestPokerHand(t *testing.T) {
	cards := "AH 2H 5H 8H JH"
	if hc, _ := PokerHand(parseCards(cards)); hc != Flush {
		t.Errorf("%#v should be flush but %s", cards, hc)
	}

	cards = "AH 2D 3C 4S 5H"
	if hc, _ := PokerHand(parseCards(cards)); hc != Straight {
		t.Errorf("%#v should be straight but %s", cards, hc)
	}

	cards = "AH 2H 3H 4H 5H"
	if hc, _ := PokerHand(parseCards(cards)); hc != StraightFlush {
		t.Errorf("%#v should be straight flush but %s", cards, hc)
	}

	cards = "AH AC 3H AS AD"
	if hc, _ := PokerHand(parseCards(cards)); hc != FourOfAKind {
		t.Errorf("%#v should be four of a kind flush but %s", cards, hc)
	}

	cards = "TH TC TD 5S 5D"
	if hc, _ := PokerHand(parseCards(cards)); hc != FullHouse {
		t.Errorf("%#v should be full house but %s", cards, hc)
	}

	cards = "3H 3C 3D JS QD"
	if hc, _ := PokerHand(parseCards(cards)); hc != ThreeOfAKind {
		t.Errorf("%#v should be three of a kind flush but %s", cards, hc)
	}

	cards = "2H 2C TD 5S 5D"
	if hc, _ := PokerHand(parseCards(cards)); hc != TwoPair {
		t.Errorf("%#v should be two pair but %s", cards, hc)
	}

	cards = "2H 2C TD 5S 3D"
	if hc, _ := PokerHand(parseCards(cards)); hc != OnePair {
		t.Errorf("%#v should be one pair but %s", cards, hc)
	}

	cards = "2H 7C TD 5S 3D"
	if hc, _ := PokerHand(parseCards(cards)); hc != HighCard {
		t.Errorf("%#v should be high card but %s", cards, hc)
	}
}
