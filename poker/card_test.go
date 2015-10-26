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

func cardsToString(cs []Card) string {
	var tmp []Card
	tmp = append(tmp, cs...)
	codes := []string{}
	for _, c := range tmp {
		codes = append(codes, c.String())
	}
	return strings.Join(codes, " ")
}

func TestCardString(t *testing.T) {
	card := Card{1, Diamonds}
	if card.String() != "AD" {
		t.Errorf("%s should be 1D.", card)
	}
	card = Card{10, Clubs}
	if card.String() != "TC" {
		t.Errorf("%s should be TC.", card)
	}
	card = Card{11, Hearts}
	if card.String() != "JH" {
		t.Errorf("%s should be JH.", card)
	}
	card = Card{12, Spades}
	if card.String() != "QS" {
		t.Errorf("%s should be QS.", card)
	}
	card = Card{13, Clubs}
	if card.String() != "KC" {
		t.Errorf("%s should be KC.", card)
	}
}

func TestParseCard(t *testing.T) {
	var card Card
	if err := ParseCard("TH", &card); err != nil {
		t.Error(err)
	}
	if card.Rank != 10 || card.Suit != Hearts {
		t.Errorf("%s should be TH", card)
	}

	if err := ParseCard("JC", &card); err != nil {
		t.Error(err)
	}
	if card.Rank != 11 || card.Suit != Clubs {
		t.Errorf("%s should be JC", card)
	}

	if err := ParseCard("QS", &card); err != nil {
		t.Error(err)
	}
	if card.Rank != 12 || card.Suit != Spades {
		t.Errorf("%s should be QS", card)
	}

	if err := ParseCard("AD", &card); err != nil {
		t.Error(err)
	}
	if card.Rank != 1 || card.Suit != Diamonds {
		t.Errorf("%s should be AD", card)
	}

	if err := ParseCard("2H", &card); err != nil {
		t.Error(err)
	}
	if card.Rank != 2 || card.Suit != Hearts {
		t.Errorf("%s should be 2H", card)
	}

	if err := ParseCard("1S", &card); err == nil {
		t.Errorf("%s should be an invalid card", card)
	}

	if err := ParseCard("11C", &card); err == nil {
		t.Errorf("%s should be an invalid card", card)
	}

	if err := ParseCard("12H", &card); err == nil {
		t.Errorf("%s should be an invalid card", card)
	}

	if err := ParseCard("13D", &card); err == nil {
		t.Errorf("%s should be an invalid card", card)
	}

	if err := ParseCard("AA", &card); err == nil {
		t.Errorf("%s should be an invalid card", card)
	}

	if err := ParseCard("HQ", &card); err == nil {
		t.Errorf("%s should be an invalid card", card)
	}
}

func TestSortCards(t *testing.T) {
	cards := parseCards("KD QH AD 5H 3D 7H KS")
	SortCards(cards)
	if cs := cardsToString(cards); cs != "AD KD KS QH 7H 5H 3D" {
		t.Errorf("%#v should be %s", cs, "AD KD KS QH 7H 5H 3D")
	}
}
