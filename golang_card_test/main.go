package main

import (
	"fmt"
	"sort"
	"strconv"
	"test/card"
)

func main() {

	//輸入區
	f2input := []int{5, 4, 3, 2, 1}
	f3input := []int{0x30a, 0x30b, 0x30c, 0x30d, 0x30e}
	f4input := []int{0x40a, 0x40b, 0x50f, 0x610, 0x20e}
	fmt.Printf("hello world: %s\n", f1())
	fmt.Printf("排序: %s\n", f2(f2input))
	fmt.Printf("算牌型: %d\n", f3(f3input))
	fmt.Printf("赖子算牌型: %d\n", f4(f4input))
}

func f1() string {
	return "hello world"
}

func f2(input []int) string {
	sort.Ints(input)
	result := ""
	for i, num := range input {
		if i > 0 {
			result += " "
		}
		result += strconv.Itoa(num)
	}
	return result
}

func f3(input []int) int {
	cardSet := new(card.CardSet)
	cardSet.Cards = input
	if len(cardSet.Cards) != 5 {
		return 0
	}
	cardSet.SetCards()
	cardSet.Sort()
	cardSet.Type = cardSet.GetSetType()
	return cardSet.Type
}

func f4(input []int) int {
	cardSet := new(card.CardSet)
	cardSet.Cards = input
	if len(cardSet.Cards) != 5 {
		return 0
	}
	cardSet.SetCards()
	sort.Ints(cardSet.Cards)
	if cardSet.Cards[3] == 0x50f && cardSet.Cards[4] == 0x610 {
		// 有小王和大王
		originalCard := make([]int, len(cardSet.Cards))
		copy(originalCard, cardSet.Cards)
		highestType, currentType := card.HighCard, card.HighCard
		for i, c := range cardSet.AllCards {
			for _, d := range cardSet.AllCards[i+1:] {
				copy(cardSet.Cards, originalCard)
				cardSet.Cards[3] = c
				cardSet.Cards[4] = d
				cardSet.Sort()
				currentType = cardSet.GetSetType()
				if currentType < highestType {
					highestType = currentType
				}
			}
		}
		cardSet.Cards = originalCard
		cardSet.Type = highestType
	} else if cardSet.Cards[4] == 0x50f || cardSet.Cards[4] == 0x610 {
		// 有小王或大王
		originalCard := make([]int, len(cardSet.Cards))
		copy(originalCard, cardSet.Cards)
		highestType, currentType := card.HighCard, card.HighCard
		for _, c := range cardSet.AllCards {
			copy(cardSet.Cards, originalCard)
			cardSet.Cards[4] = c
			cardSet.Sort()
			currentType = cardSet.GetSetType()
			if currentType < highestType {
				highestType = currentType
			}
		}
		cardSet.Cards = originalCard
		cardSet.Type = highestType
	} else {
		cardSet.Sort()
		cardSet.Type = cardSet.GetSetType()
	}
	return cardSet.Type
}
