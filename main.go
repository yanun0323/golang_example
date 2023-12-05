package main

import (
	"fmt"
	"main/game"
)

var (
	//張數
	max = 52
	//牌組數
	group = 1
	//玩家數量
	playerCnt = 2
	//玩家牌數
	cardCnt = 3
)

func main() {
	//初始牌靴
	table, err := game.InitTable(max, group)
	if err != nil {
		fmt.Println(err)
		return
	}

	drawed, err1 := game.Draw(table, cardCnt)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(drawed)

	//比大小
	max, err2 := game.Compare(drawed)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(max)

}
