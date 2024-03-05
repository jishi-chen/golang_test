package card

import (
	"slices"
	"sort"
)

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
	Nine  = 0x09
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
	AllCards []int
	Type     int
	Cards    []int
}

// 交錯產生所有牌組
func (s *CardSet) SetCards() {
	var cards []int
	for i := Two; i <= Ace; i++ {
		cards = append(cards, Diamonds|i)
		cards = append(cards, Clubs|i)
		cards = append(cards, Hearts|i)
		cards = append(cards, Spades|i)
	}
	for _, card := range cards {
		if !slices.Contains(s.Cards, card) {
			s.AllCards = append(s.AllCards, card)
		}
	}
}

// 排序
func (s *CardSet) Sort() {
	sort.Ints(s.Cards) //排序花色
	sort.Slice(s.Cards, func(i, j int) bool {
		return s.Cards[i]&0x0ff < s.Cards[j]&0x0ff // 排序數字大小
	})
}

// 是否為同花順
func (s *CardSet) isStraightFlush() bool {
	return s.isStraight() && s.isFlush()
}

// 是否為同花
func (s *CardSet) isFlush() bool {
	suits := make(map[int]bool)
	for _, c := range s.Cards {
		suits[c&0xf00] = true
	}
	return len(suits) == 1
}

// 是否為順子
func (s *CardSet) isStraight() bool {
	for i := 1; i < len(s.Cards); i++ {
		if s.Cards[i]-s.Cards[i-1] != 1 {
			return false
		}
	}
	return true
}

// 是否為金剛
func (s *CardSet) isFourOfAKind() bool {
	return s.Cards[0]&0x00F == s.Cards[3]&0x00F || s.Cards[1]&0x00F == s.Cards[4]&0x00F
}

// 是否為葫蘆
func (s *CardSet) isFullHouse() bool {
	return (s.Cards[0]&0x00F == s.Cards[1]&0x00F && s.Cards[2]&0x00F == s.Cards[4]&0x00F) || (s.Cards[0]&0x00F == s.Cards[2]&0x00F && s.Cards[3]&0x00F == s.Cards[4]&0x00F)
}

// 是否為三條
func (s *CardSet) isThreeOfAKind() bool {
	return (s.Cards[0]&0x00F == s.Cards[2]&0x00F) || (s.Cards[1]&0x00F == s.Cards[3]&0x00F) || (s.Cards[2]&0x00F == s.Cards[4]&0x00F)
}

// 是否為兩對
func (s *CardSet) isTwoPairs() bool {
	if s.Cards[0]&0x00F == s.Cards[1]&0x00F {
		return s.Cards[2]&0x00F == s.Cards[3]&0x00F || s.Cards[3]&0x00F == s.Cards[4]&0x00F
	}
	if s.Cards[1]&0x00F == s.Cards[2]&0x00F {
		return s.Cards[3]&0x00F == s.Cards[4]&0x00F
	}
	return false
}

// 是否為一對
func (s *CardSet) isOnePair() bool {
	return (s.Cards[0]&0x00F == s.Cards[1]&0x00F) || (s.Cards[1]&0x00F == s.Cards[2]&0x00F) || (s.Cards[2]&0x00F == s.Cards[3]&0x00F) || (s.Cards[3]&0x00F == s.Cards[4]&0x00F)
}

// 是否為皇家同花順
func (s *CardSet) isRoyalFlush() bool {
	return s.isFlush() &&
		s.Cards[0]&0x00F == 10 && s.Cards[1]&0x00F == 11 && s.Cards[2]&0x00F == 12 && s.Cards[3]&0x00F == 13 && s.Cards[4]&0x00F == 14
}

// 判断牌型
func (s *CardSet) GetSetType() int {

	switch {
	case s.isRoyalFlush():
		return RoyalFlush
	case s.isStraightFlush():
		return StraightFlush
	case s.isFourOfAKind():
		return FourOfAKind
	case s.isFullHouse():
		return FullHouse
	case s.isFlush():
		return Flush
	case s.isStraight():
		return Straight
	case s.isThreeOfAKind():
		return ThreeOfAKind
	case s.isTwoPairs():
		return TwoPairs
	case s.isOnePair():
		return OnePair
	default:
		return HighCard
	}
}
