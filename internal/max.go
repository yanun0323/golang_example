package internal

import "errors"

// maxWithLimitation 比大小，如果一樣大回傳 {{0, 0}, nil}
func MaxWithLimitation(a, b [][]int) ([][]int, error) {

	sumScoreA, errA := SumScore(a)

	if errA != nil {
		return nil, errA
	}

	sumScoreB, errB := SumScore(b)

	if errB != nil {
		return nil, errB
	}

	for V := range a {
		if a[V][0] < 1 || a[V][0] > 13 {
			return nil, errors.New("a < 1 || a > 13")
		}
	}

	for V := range b {
		if b[V][0] < 1 || b[V][0] > 13 {
			return nil, errors.New("b < 1 || b > 13")
		}
	}

	if sumScoreA > sumScoreB {
		return a, nil
	} else if sumScoreA < sumScoreB {
		return b, nil
	}

	sumSuitA, errSA := SumSuit(a)

	if errSA != nil {
		return nil, errSA
	}

	sumSuitB, errSB := SumSuit(b)

	if errSB != nil {
		return nil, errSB
	}

	for V := range a {
		if a[V][1] < 0 || a[V][1] > 3 {
			return nil, errors.New("a < 0 || a > 3")
		}
	}

	for V := range b {
		if b[V][1] < 0 || b[V][1] > 3 {
			return nil, errors.New("b < 0 || b > 3")
		}
	}

	if sumSuitA > sumSuitB {
		return a, nil
	} else if sumSuitA < sumSuitB {
		return b, nil
	}

	return [][]int{{0, 0}, nil}, nil

}
