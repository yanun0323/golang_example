package algorithm

import (
	"errors"
	"main/internal"
)

type step2 struct {
	cards [][]int
}

func NewStep2(cards [][]int) Algorithm {
	return &step2{
		cards: cards,
	}
}

func (s *step2) CalculateAndGetWinner() ([]int, error) {
	if len(s.cards) <= 0 {
		return nil, errors.New("cards should greater than 0")
	}

	max := s.cards[0]
	for i := 0; i < len(s.cards); i++ {
		tempMax, err := internal.MaxWithLimitation(max, s.cards[i])
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
