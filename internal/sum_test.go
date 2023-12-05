package internal

import "testing"

func TestSumScore(t *testing.T) {

	tests := []struct {
		name    string
		cards   [][]int
		want    int
		wantErr bool
	}{
		{"good", [][]int{{1, 0}, {2, 3}, {3, 1}}, 6, false},
		{"cards != 3", [][]int{{2, 3}, {3, 1}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumScore(tt.cards)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumSuit(t *testing.T) {

	tests := []struct {
		name    string
		cards   [][]int
		want    int
		wantErr bool
	}{
		{"good", [][]int{{1, 0}, {2, 3}, {3, 1}}, 4, false},
		{"cards != 3", [][]int{{2, 3}, {3, 1}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumSuit(tt.cards)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumSuit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SumSuit() = %v, want %v", got, tt.want)
			}
		})
	}
}
