package game

import "errors"

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

// Draw
func Draw() {

}

// Compare
func Compare() {

}
