package game

import "testing"

func TestInitTable(t *testing.T) {
	tests := []struct {
		name       string
		max, group int
		expected   []int
		err        bool
	}{
		{"good", 3, 1, []int{1, 2, 3}, false},
		{"max < 0", -1, 1, nil, true},
		{"group < 0", 1, -1, nil, true},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table, err := InitTable(tt.max, tt.group)
			if tt.err {
				if err == nil {
					t.Fatal("nil error")
				}
				return
			}

			if err != nil {
				t.Fatal("error")
			}

			for i := range tt.expected {
				if table[i] != tt.expected[i] {
					t.Fatalf("mismatch: %d, %d", table[i], tt.expected[i])
				}
			}
		})
	}
}

func TestDraw(t *testing.T) {

	// playerCard, err := Draw(table, playerCount)

	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func TestCompare(t *testing.T) {
	// winIndex, err := Compare(playerCard)

	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Compare()
		})
	}
}
