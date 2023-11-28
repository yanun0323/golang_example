package game

import (
	"errors"
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
// 如果抽到的牌是-1 代表已經抽過了, 回傳-1
func Draw(cards []int) int {
	//產生一個範圍內整數隨機數
	r := rand.Intn(len(cards))
	cardDraw := cards[r]

	if cardDraw == -1 {
		return -1
	}
	cards[r] = -1
	return cardDraw
}

// Compare
func Compare() {
	//比較牌的點數大小
	//用maxWithLimitation

}
