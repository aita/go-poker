package poker

import "testing"

func TestPokerHand(t *testing.T) {
	cards := "AH AS 5D 8H JH QH 9H"
	ph := NewPokerHand(parseCards(cards))
	if ph.HandCategory != Flush {
		t.Errorf("%#v should be flush but %s", ph.Cards, ph.HandCategory)
	}
	if cs := cardsToString(ph.Cards); cs != "AH QH JH 9H 8H" {
		t.Errorf("%#v should be %s", cs, "AH QH JH 9H 8H")
	}

	cards = "AH 2D 3C 6H 8D 4S 5H"
	ph = NewPokerHand(parseCards(cards))
	if ph.HandCategory != Straight {
		t.Errorf("%#v should be straight but %s", cards, ph.HandCategory)
	}
	if cs := cardsToString(ph.Cards); cs != "6H 5H 4S 3C 2D" {
		t.Errorf("%#v should be %s", cs, "6H 5H 4S 3C 2D")
	}

	cards = "AH 2H JD QS 3H 4H 5H"
	ph = NewPokerHand(parseCards(cards))
	if ph.HandCategory != StraightFlush {
		t.Errorf("%#v should be straight flush but %s", cards, ph.HandCategory)
	}
	if cs := cardsToString(ph.Cards); cs != "AH 5H 4H 3H 2H" {
		t.Errorf("%#v should be %s", cs, "AH 5H 4H 3H 2H")
	}

	cards = "AH AC 3H AS TS AD"
	ph = NewPokerHand(parseCards(cards))
	if ph.HandCategory != FourOfAKind {
		t.Errorf("%#v should be four of a kind flush but %s", cards, ph.HandCategory)
	}
	if cs := cardsToString(ph.Cards); cs != "AH AC AS AD" {
		t.Errorf("%#v should be %s", cs, "AH AC AS AD")
	}

	cards = "TH 2S TC 4D TD 5S 5D"
	ph = NewPokerHand(parseCards(cards))
	if ph.HandCategory != FullHouse {
		t.Errorf("%#v should be full house but %s", cards, ph.HandCategory)
	}
	if cs := cardsToString(ph.Cards); cs != "TH TC TD 5S 5D" {
		t.Errorf("%#v should be %s", cs, "TH TC TD 5S 5D")
	}

	cards = "3H 3C 3D JS QD"
	ph = NewPokerHand(parseCards(cards))
	if ph.HandCategory != ThreeOfAKind {
		t.Errorf("%#v should be three of a kind flush but %s", cards, ph.HandCategory)
	}
	if cs := cardsToString(ph.Cards); cs != "3H 3C 3D" {
		t.Errorf("%#v should be %s", cs, "3H 3C 3D")
	}

	cards = "2H 2C TS 5S 5D TD"
	ph = NewPokerHand(parseCards(cards))
	if ph.HandCategory != TwoPair {
		t.Errorf("%#v should be two pair but %s", cards, ph.HandCategory)
	}
	if cs := cardsToString(ph.Cards); cs != "TS TD 5S 5D" {
		t.Errorf("%#v should be %s", cs, "TS TD 5S 5D")
	}

	cards = "JH 2C TD JD 5S 3D"
	ph = NewPokerHand(parseCards(cards))
	if ph.HandCategory != OnePair {
		t.Errorf("%#v should be one pair but %s", cards, ph.HandCategory)
	}
	if cs := cardsToString(ph.Cards); cs != "JH JD" {
		t.Errorf("%#v should be %s", cs, "JH JD")
	}

	cards = "2H 7C TD 5S 3D"
	ph = NewPokerHand(parseCards(cards))
	if ph.HandCategory != HighCard {
		t.Errorf("%#v should be high card but %s", cards, ph.HandCategory)
	}
	if cs := cardsToString(ph.Cards); cs != "TD" {
		t.Errorf("%#v should be %s", cs, "TD")
	}
}
