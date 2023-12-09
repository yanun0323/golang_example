package main

import (
	"fmt"
	"main/game"
	"main/game/algorithm"
)

var (
	//張數
	max = 52
	//牌組數
	group = 2
	//玩家數量
	playerCnt = 3
)

func main() {
	//初始牌靴
	table, err := game.InitTable(max, group)
	if err != nil {
		fmt.Println(err)
		return
	}

	drawed, err1 := game.Draw(table, playerCnt)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(drawed)

	alg := algorithm.NewStep2(drawed)
	_ = algorithm.NewStep1()

	//比大小
	max, err2 := game.Compare(alg)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(max)

}

func CalExample() {
	m := [][]byte{}

	p := 0
	b := 0

	pBit := 0
	p += drawCard()
	b += drawCard()

	for {
		switch m[p][b] {
		case 'H':
			p += drawCard()
		case 'S':
			break
		case 'D':
			pBit *= 2
			p += drawCard()
		case 'L':
			break
		}
	}

}

func drawCard() int {
	return 1
}
