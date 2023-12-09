package game

import (
	"errors"
	"main/game/algorithm"
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
// table: 牌靴, playerCnt: 玩家數量
func Draw(table []int, playerCnt int) ([][]int, error) {

	if len(table) < playerCnt {
		return nil, errors.New("table is not enough")
	}

	if playerCnt <= 0 {
		return nil, errors.New("playerCnt should greater than 0")
	}

	//num:存放抽到的牌
	playerNum := make([][]int, 0, playerCnt)

	for len(playerNum) < playerCnt {
		r := rand.Intn(len(table))
		if table[r] == -1 { //處理重複抽
			continue
		}

		num := table[r]%13 + 1
		suit := table[r] / 13

		playerNum = append(playerNum, []int{num, suit})
		table[r] = -1
	}

	if len(playerNum) != playerCnt {
		return nil, errors.New("card num is not match")
	}

	return playerNum, nil
}

// Compare: 比大小 傳回最大值 有錯傳回-1及錯誤資訊
func Compare(alg algorithm.Algorithm) ([]int, error) {
	return alg.CalculateAndGetWinner()

}

// DrawCard 抽牌示範
func DrawCard(count int) []int {
	cards := []int{}
	p := 0
	picked := map[int]bool{}
	pick := make([]int, 0, count)
	for i := 0; i < count; {
		p = rand.Intn(52)
		if picked[p] {
			continue
		}
		pick = append(pick, cards[p])
		picked[p] = true
		i++
	}
	return pick
}
