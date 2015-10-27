package poker

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

type PokerHand struct {
	HandCategory HandCategory
	Cards        []Card
}

func NewPokerHand(cards []Card) *PokerHand {
	ph := PokerHand{}
	SortCards(cards)
	if ph2 := checkStraight(cards); ph2 != nil {
		ph = *ph2
	}
	if ph2 := checkPairs(cards); ph.HandCategory <= ph2.HandCategory {
		ph = *ph2
	}
	if ph.HandCategory < Flush {
		if ph2 := checkFlush(cards); ph2 != nil {
			ph = *ph2
		}
	}
	return &ph
}

func checkStraight(cards []Card) *PokerHand {
	ph := PokerHand{}
	bins := [13][]Card{}
	for _, c := range cards {
		bins[c.Rank-1] = append(bins[c.Rank-1], c)
	}

	var n int
	ranks := [...]Rank{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 1}
	for i, rank := range ranks {
		if bins[rank-1] != nil {
			n++
		} else {
			n = 0
		}
		if n >= 5 {
			cs := []Card{}
			for _, r := range ranks[i-4 : i+1] {
				cs = append(cs, bins[r-1]...)
			}
			if ph2 := checkFlush(cs); ph2 != nil {
				ph.Cards = ph2.Cards
				ph.HandCategory = StraightFlush
			} else if ph.HandCategory != StraightFlush {
				ph.Cards = cs
				ph.HandCategory = Straight
			}
		}
	}
	if ph.HandCategory == Straight || ph.HandCategory == StraightFlush {
		SortCards(ph.Cards)
		if len(ph.Cards) > 5 {
			ph.Cards = ph.Cards[:5]
		}
		return &ph
	}
	return nil
}

func checkFlush(cards []Card) *PokerHand {
	ph := PokerHand{}
	bins := [4][]Card{}
	for _, c := range cards {
		bins[c.Suit] = append(bins[c.Suit], c)
	}

	for _, cs := range bins {
		if len(cs) < 5 {
			continue
		}
		ph.HandCategory = Flush
		for _, card := range cards {
			for _, c := range cs {
				if card == c {
					ph.Cards = append(ph.Cards, c)
				}
			}
		}
		if len(ph.Cards) > 5 {
			ph.Cards = ph.Cards[:5]
		}
		return &ph
	}
	return nil
}

func checkPairs(cards []Card) *PokerHand {
	ph := PokerHand{}
	bins := [13][]Card{}
	for _, c := range cards {
		bins[c.Rank-1] = append(bins[c.Rank-1], c)
	}

	two := [13]int{}
	three := [13]int{}
	four := [13]int{}
	for i, cs := range bins {
		switch len(cs) {
		case 2:
			two[i]++
		case 3:
			three[i]++
		case 4:
			four[i]++
		}
	}
	for i, n := range four {
		if n > 0 {
			ph.Cards = bins[i]
			ph.HandCategory = FourOfAKind
			return &ph
		}
	}

	for i, n := range three {
		if n > 0 {
			ph.HandCategory = ThreeOfAKind
			ph.Cards = bins[i]
		}
	}
	for _, rank := range [...]Rank{1, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2} {
		if two[rank-1] > 0 {
			if ph.HandCategory == ThreeOfAKind {
				ph.HandCategory = FullHouse
				ph.Cards = append(ph.Cards, bins[rank-1]...)
				return &ph
			} else if ph.HandCategory == OnePair {
				ph.HandCategory = TwoPair
				ph.Cards = append(ph.Cards, bins[rank-1]...)
				return &ph
			} else {
				ph.HandCategory = OnePair
				ph.Cards = append(ph.Cards, bins[rank-1]...)
			}
		}
	}
	if ph.HandCategory == HighCard {
		ph.HandCategory = HighCard
		ph.Cards = cards[:1]
	}
	return &ph
}
