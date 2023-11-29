package internal

import "errors"

// maxWithLimitation 比大小，如果一樣大回傳 {-1,-1}
func MaxWithLimitation(a, b []int) ([]int, error) {
	if a[0] > 13 {
		return nil, errors.New("a > 13")
	}

	if b[0] > 13 {
		return nil, errors.New("b > 13")
	}

	if a[0] <= 0 {
		return nil, errors.New("a <= 0")
	}

	if b[0] <= 0 {
		return nil, errors.New("b <= 0")
	}

	if a[1] < 0 {
		return nil, errors.New("a[1] < 0")
	}

	if b[1] < 0 {
		return nil, errors.New("b[1] < 0")
	}

	if a[1] > 3 {
		return nil, errors.New("a[1] > 3")
	}

	if b[1] > 3 {
		return nil, errors.New("b[1] > 3")
	}

	if a[0] > b[0] {
		return a, nil
	}

	if a[0] == b[0] { //比較花色
		if a[1] > b[1] {
			return a, nil
		}

		if a[1] < b[1] {
			return b, nil
		}

		return []int{-1, -1}, nil
	}

	return b, nil

}
