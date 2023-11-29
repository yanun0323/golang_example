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
			cards = append(cards, j+1)
		}
	}
	return cards, nil
}

// Draw: Draw a card from table
// table: 牌靴, playerCnt: 玩家數量
func Draw(table []int, playerCnt int) ([]int, error) {

	if len(table) < playerCnt {
		return nil, errors.New("table is not enough")
	}

	if playerCnt <= 0 {
		return nil, errors.New("playerCnt should greater than 0")
	}

	//num:存放抽到的牌
	playerNum := make([]int, 0, playerCnt)

	for len(playerNum) < playerCnt {
		r := rand.Intn(len(table))
		if table[r] == -1 { //處理重複抽
			continue
		}
		playerNum = append(playerNum, table[r])
		table[r] = -1
	}

	if len(playerNum) != playerCnt {
		return nil, errors.New("card num is not match")
	}

	return playerNum, nil
}

// Compare: 比大小 傳回最大值 有錯傳回-1及錯誤資訊
func Compare(cards []int) (int, error) {
	if len(cards) <= 0 {
		return -1, errors.New("cards should greater than 0")
	}

	max := 0
	for i := 0; i < len(cards); i++ {
		tempMax, err := internal.MaxWithLimitation(max, cards[i])
		if err == nil {
			if tempMax != -1 {
				max = tempMax
			}
		} else {
			return -1, err
		}
	}

	return max, nil
}
