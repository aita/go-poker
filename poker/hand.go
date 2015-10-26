package poker

import "log"

type HandCategory int

const (
	HighCard = HandCategory(iota)
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

func (hc HandCategory) String() string {
	switch hc {
	case HighCard:
		return "high card"
	case OnePair:
		return "one pair"
	case TwoPair:
		return "two pair"
	case ThreeOfAKind:
		return "three of a kind"
	case Straight:
		return "straight"
	case Flush:
		return "flush"
	case FullHouse:
		return "full house"
	case FourOfAKind:
		return "four of a kind"
	case StraightFlush:
		return "straight flush"
	}
	return ""
}

var straightRanks = []Rank{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 1}

func checkStaright(cards []Card) []Card {
	ranks := map[Rank][]Card{}
	for _, c := range cards {
		ranks[c.Rank] = append(ranks[c.Rank], c)
	}

	var high Rank
	n, pos := 0, 0
	for i, rank := range straightRanks {
		if _, ok := ranks[rank]; ok {
			n++
		} else {
			n = 0
		}
		if n >= 5 {
			high = rank
			pos = i
		}
	}
	var result []Card
	if high > 0 {
		for _, rank := range straightRanks[pos-4 : pos+1] {
			cs, _ := ranks[rank]
			result = append(result, cs...)
		}
		SortCards(result)
	}
	return result
}

func checkFlush(cards []Card) []Card {
	suits := map[Suit][]Card{}
	for _, c := range cards {
		suits[c.Suit] = append(suits[c.Suit], c)
	}

	for _, cs := range suits {
		if len(cs) >= 5 {
			SortCards(cs)
			return cs[:5]
		}
	}
	return nil
}

func checkPairs(cards []Card) (HandCategory, []Card) {
	ranks := map[Rank][]Card{}
	for _, c := range cards {
		ranks[c.Rank] = append(ranks[c.Rank], c)
	}

	two := [13]int{}
	three := [13]int{}
	four := [13]int{}
	for rank, cs := range ranks {
		switch len(cs) {
		case 2:
			two[rank]++
		case 3:
			three[rank]++
		case 4:
			four[rank]++
		}
	}
	for rank, n := range four {
		if n > 0 {
			return FourOfAKind, ranks[Rank(rank)]
		}
	}
	hc := HighCard
	result := []Card{}
	var high Rank
	for rank, n := range three {
		if n > 0 {
			hc = ThreeOfAKind
			if high < Rank(rank) {
				high = Rank(rank)
				result = ranks[Rank(rank)]
			}
		}
	}
	for _, rank := range []Rank{1, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2} {
		if two[rank-1] > 0 {
			if hc == ThreeOfAKind {
				hc = FullHouse
				result = append(result, ranks[rank-1]...)
				return hc, result
			} else if hc == OnePair {
				hc = TwoPair
				result = append(result, ranks[rank-1]...)
				return hc, result
			} else {
				hc = OnePair
				result = append(result, ranks[rank-1]...)
			}
		}
	}
	if hc == HighCard {
		result = cards[:1]
		log.Println(cards, result)
	}
	return hc, result
}

type PokerHand struct {
	HandCategory HandCategory
	Cards        []Card
}

func NewPokerHand(cards []Card) *PokerHand {
	ph := PokerHand{}
	SortCards(cards)
	if cs := checkStaright(cards); cs != nil {
		if cs2 := checkFlush(cs); cs2 != nil {
			ph.HandCategory = StraightFlush
			ph.Cards = cs2
			return &ph
		}
		ph.HandCategory = Straight
		ph.Cards = cs
	}
	if cs := checkFlush(cards); cs != nil {
		ph.HandCategory = Flush
		ph.Cards = cs
	}
	if hc, cs := checkPairs(cards); ph.HandCategory <= hc {
		ph.HandCategory = hc
		ph.Cards = cs
	}
	return &ph
}
