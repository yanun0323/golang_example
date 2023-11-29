package main

import (
	"fmt"
	"main/game"
)

func main() {
	//初始牌靴
	table, err := game.InitTable(52, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	drawed, err1 := game.Draw(table, 3)
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
