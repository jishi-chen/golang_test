package mymath

import "sort"

const (
	Spades   = 0x400
	Hearts   = 0x300
	Clubs    = 0x200
	Diamonds = 0x100
)

const (
	Two   = 0x02
	Three = 0x03
	Four  = 0x04
	Five  = 0x05
	Six   = 0x06
	Seven = 0x07
	Eight = 0x08
	Nine  = 0x0910
	Ten   = 0x0a
	Jack  = 0x0b
	Queen = 0x0c
	King  = 0x0d
	Ace   = 0x0e
)

const (
	RoyalFlush    = 1
	StraightFlush = 2
	FourOfAKind   = 3
	FullHouse     = 4
	Flush         = 5
	Straight      = 6
	ThreeOfAKind  = 7
	TwoPairs      = 8
	OnePair       = 9
	HighCard      = 10
)

type CardSet struct {
	Type  int
	Cards []int
}

// 排序
func (s *CardSet) Sort() {
	sort.Ints(s.Cards)
}

// 是否為同花順
func (s *CardSet) isStraightFlush() bool {
	return s.isStraight() && s.isFlush()
}

// 是否為同花
func (s *CardSet) isFlush() bool {
	if len(s.Cards) != 5 {
		return false
	}
	suits := make(map[int]bool)
	for _, c := range s.Cards {
		suits[c&0xf00] = true
	}
	return len(suits) == 1
}

// 是否為順子
func (s *CardSet) isStraight() bool {
	if len(s.Cards) != 5 {
		return false
	}
	for i := 1; i < len(s.Cards); i++ {
		if s.Cards[i]-s.Cards[i-1] != 1 {
			return false
		}
	}
	return true
}

// 是否為金剛
func (s *CardSet) isFourOfAKind() bool {
	if len(s.Cards) != 5 {
		return false
	}
	return s.Cards[0] == s.Cards[3] || s.Cards[1] == s.Cards[4]
}

// 是否為葫蘆
func (s *CardSet) isFullHouse() bool {
	if len(s.Cards) != 5 {
		return false
	}
	return (s.Cards[0] == s.Cards[1] && s.Cards[2] == s.Cards[4]) || (s.Cards[0] == s.Cards[2] && s.Cards[3] == s.Cards[4])
}

// 是否為三條
func (s *CardSet) isThreeOfAKind() bool {
	if len(s.Cards) != 5 {
		return false
	}
	return (s.Cards[0] == s.Cards[2]) || (s.Cards[1] == s.Cards[3]) || (s.Cards[2] == s.Cards[4])
}

// 是否為兩對
func (s *CardSet) isTwoPairs() bool {
	if len(s.Cards) != 5 {
		return false
	}
	if s.Cards[0] == s.Cards[1] {
		return s.Cards[2] == s.Cards[3] || s.Cards[3] == s.Cards[4]
	}
	if s.Cards[1] == s.Cards[2] {
		return s.Cards[3] == s.Cards[4]
	}
	return false
}

// 是否為一對
func (s *CardSet) isOnePair() bool {
	if len(s.Cards) != 5 {
		return false
	}
	return (s.Cards[0] == s.Cards[1]) || (s.Cards[1] == s.Cards[2]) || (s.Cards[2] == s.Cards[3]) || (s.Cards[3] == s.Cards[4])
}

// 是否為皇家同花順
func (s *CardSet) isRoyalFlush() bool {
	if len(s.Cards) != 5 {
		return false
	}
	return s.Cards[0] == 0x10a && s.Cards[1] == 0x10b && s.Cards[2] == 0x10c && s.Cards[3] == 0x10d && s.Cards[4] == 0x10e
}

// 判断牌型
func (s *CardSet) GetSetType() {
	switch {
	case s.isRoyalFlush():
		s.Type = RoyalFlush
	case s.isStraightFlush():
		s.Type = StraightFlush
	case s.isFourOfAKind():
		s.Type = FourOfAKind
	case s.isFullHouse():
		s.Type = FullHouse
	case s.isFlush():
		s.Type = Flush
	case s.isStraight():
		s.Type = Straight
	case s.isThreeOfAKind():
		s.Type = ThreeOfAKind
	case s.isTwoPairs():
		s.Type = TwoPairs
	case s.isOnePair():
		s.Type = OnePair
	default:
		s.Type = HighCard
	}
}
