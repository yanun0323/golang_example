package game

import (
	"errors"
	"main/internal"
	"math/rand"
)

// InitTable
func InitTable(max, group int) ([]int, error) {
	if max <= 0 {
		return nil, errors.New("max should greater than 0")
	}

	if group <= 0 {
		return nil, errors.New("group should greater than 0")
	}

	cards := make([]int, 0, max*group)
	for i := 0; i < group; i++ {
		for j := 0; j < max; j++ {
			cards = append(cards, j)
		}
	}
	return cards, nil
}

// Draw: Draw a card from table
// table: 牌靴, cardsPerP: 每個玩家的牌數
// 回傳: 每個玩家的牌{點數, 花色}, error
func Draw(table []int, cardsPerP int) ([][]int, error) {

	if len(table) < cardsPerP {
		return nil, errors.New("table is not enough")
	}

	if cardsPerP <= 0 {
		return nil, errors.New("cardsPerP should greater than 0")
	}

	//num:存放抽到的牌
	playerCard := make([][]int, 0, cardsPerP)

	for len(playerCard) < cardsPerP {
		r := rand.Intn(len(table))
		if table[r] == -1 { //處理重複抽
			continue
		}

		num := table[r]%13 + 1
		suit := table[r] / 13

		playerCard = append(playerCard, []int{num, suit})
		table[r] = -1
	}

	if len(playerCard) != cardsPerP {
		return nil, errors.New("card num is not match")
	}

	return playerCard, nil
}

// Compare: 比大小 傳回最大值 有錯傳回-1及錯誤資訊
func Compare(p1, p2 [][]int) ([][]int, error) {
	if len(p1) <= 0 {
		return nil, errors.New("p1 cards should greater than 0")
	}

	if len(p2) <= 0 {
		return nil, errors.New("p2 cards should greater than 0")
	}

	if err != nil {
		return nil, err
	}

	internal.MaxWithLimitation(scoreSumP1, scoreSumP2)

	if scoreSumP1 > scoreSumP2 {
		return p1, nil
	} else if scoreSumP1 < scoreSumP2 {
		return p2, nil
	}

	suitSumP1, err := internal.SumSuit(p1)
	suitSumP2, err := internal.SumSuit(p2)

	if err != nil {
		return nil, err
	}

	if suitSumP1 > suitSumP2 {
		return p1, nil
	} else if suitSumP1 < suitSumP2 {
		return p2, nil
	} else {
		return [][]int{-1, -1}, nil
	}

	max := cards[0]
	for i := 0; i < len(cards); i++ {
		tempMax, err := internal.MaxWithLimitation(max, cards[i])
		if err == nil {
			if tempMax[0] != -1 && tempMax[1] != -1 {
				max = tempMax
			}
		} else {
			return nil, err
		}
	}

	return max, nil
}
