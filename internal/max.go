package internal

import "errors"

// maxWithLimitation 比大小，如果一樣大回傳 -1
func maxWithLimitation(a, b int) (int, error) {
	if a > 52 {
		return 0, errors.New("a can not greater than 52")
	}

	if b > 52 {
		return 0, errors.New("b can not greater than 52")
	}

	if a < 0 {
		return 0, errors.New("a can not less than 0")
	}

	if b < 0 {
		return 0, errors.New("b can not less than 0")
	}

	if a > b {
		return a, nil
	}

	if a == b {
		return -1, nil
	}

	return b, nil
}
