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

	tests := []struct {
		name        string
		table       []int
		playerCount int
		expectedLen int
		err         bool
	}{
		{"good", []int{1, 2, 3}, 2, 2, false},
		{"good", []int{1, 2, 3}, 3, 3, false},
		{"card < player", []int{3}, 2, 0, true},
		{"card < player", nil, 2, 0, true},
		{"player <= 0", []int{1, 2, 3}, 0, 0, true},
		{"player <= 0", []int{1, 2, 3}, -1, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			draw, err := Draw(tt.table, tt.playerCount)
			if tt.err {
				if err == nil {
					t.Fatal("nil error")
				}
				return
			}

			if err != nil {
				t.Fatal(err)
			}

			if len(draw) != tt.expectedLen {
				t.Fatalf("mismatch: %d, %d", len(draw), tt.expectedLen)
			}

		})
	}
}

func TestCompare(t *testing.T) {
	// winIndex, err := Compare(playerCard)

	tests := []struct {
		name     string
		cards    []int
		expected int
		err      bool
	}{
		{"good", []int{1, 2, 3}, 3, false},
		{"good", []int{2, 2, 3, 3}, 3, false},
		{"good", []int{1, 2, 4, 4}, 4, false},
		{"good", []int{2, 2, 2, 2, 2}, 2, false},
		{"cards <= 0", nil, -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			com, err := Compare(tt.cards)
			if tt.err {
				if err == nil {
					t.Fatal("nil error")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}

			if com != tt.expected {
				t.Fatalf("mismatch: %d, %d", com, tt.expected)

			}
		})
	}
}
