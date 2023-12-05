package internal

import "errors"

func SumScore(card [][]int) (int, error) {
	if len(card) != 3 {
		return 0, errors.New("card num is not match")
	}

	sum := 0
	for i := 0; i < len(card); i++ {
		sum += card[i][0]
	}

	return sum, nil
}

func SumSuit(card [][]int) (int, error) {
	if len(card) != 3 {
		return 0, errors.New("card num is not match")
	}

	sum := 0
	for i := 0; i < len(card); i++ {
		sum += card[i][1]
	}

	return sum, nil

}
