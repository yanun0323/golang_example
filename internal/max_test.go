package internal

import (
	"testing"
)

func TestMaxWithLimit(t *testing.T) {
	testCase := []struct {
		name     string
		a, b     [][]int
		expected [][]int
		isError  bool
	}{
		{
			name:     "a 比較大",
			a:        [][]int{{5, 0}, {4, 1}, {13, 1}},
			b:        [][]int{{4, 1}, {3, 2}, {12, 1}},
			expected: [][]int{{5, 0}, {4, 1}, {13, 1}},
		},
		{
			name:     "b 比較大",
			a:        [][]int{{5, 0}, {4, 1}, {13, 1}},
			b:        [][]int{{6, 1}, {6, 2}, {12, 1}},
			expected: [][]int{{6, 1}, {6, 2}, {12, 1}},
		},
		{
			name:     "a 比較大(花色)",
			a:        [][]int{{5, 3}, {4, 1}, {13, 1}},
			b:        [][]int{{4, 1}, {5, 2}, {13, 1}},
			expected: [][]int{{5, 3}, {4, 1}, {13, 1}},
		},
		{
			name:     "b 比較大(花色)",
			a:        [][]int{{5, 1}, {4, 1}, {13, 1}},
			b:        [][]int{{6, 3}, {5, 2}, {11, 1}},
			expected: [][]int{{6, 3}, {5, 2}, {11, 1}},
		},
		{
			name:     "a 數字超過",
			a:        [][]int{{15, 1}, {4, 1}, {13, 1}},
			b:        [][]int{{4, 3}, {5, 2}, {13, 1}},
			expected: nil,
			isError:  true,
		},
		{
			name:     "b 數字超過",
			a:        [][]int{{5, 1}, {4, 1}, {13, 1}},
			b:        [][]int{{14, 3}, {5, 2}, {13, 1}},
			expected: nil,
			isError:  true,
		},
		{
			name:     "a < lowerLimit",
			a:        [][]int{{0, 1}, {1, 1}, {13, 1}},
			b:        [][]int{{4, 3}, {5, 2}, {13, 1}},
			expected: nil,
			isError:  true,
		},
		{
			name:     "b < lowerLimit",
			a:        [][]int{{10, 1}, {1, 1}, {13, 1}},
			b:        [][]int{{1, 3}, {-2, 2}, {13, 1}},
			expected: nil,
			isError:  true,
		},
		{
			name:     "a 花色<0",
			a:        [][]int{{10, 1}, {1, -1}, {13, 1}},
			b:        [][]int{{11, 3}, {2, 2}, {13, 1}},
			expected: nil,
			isError:  true,
		},
		{
			name:     "b 花色<0",
			a:        [][]int{{10, 1}, {1, 1}, {13, 1}},
			b:        [][]int{{11, 3}, {2, -2}, {13, 1}},
			expected: nil,
			isError:  true,
		},
		{
			name:     "a 花色超過",
			a:        [][]int{{10, 1}, {1, 4}, {13, 1}},
			b:        [][]int{{11, 3}, {2, 2}, {13, 1}},
			expected: nil,
			isError:  true,
		},
		{
			name:     "b 花色超過",
			a:        [][]int{{10, 1}, {1, 2}, {13, 1}},
			b:        [][]int{{11, 5}, {2, 2}, {13, 1}},
			expected: nil,
			isError:  true,
		},
		{
			name:     "平手",
			a:        [][]int{{3, 1}, {2, 1}, {13, 1}},
			b:        [][]int{{4, 1}, {1, 1}, {13, 1}},
			expected: [][]int{{0, 0}, nil},
			isError:  false,
		},
	}

	for _, tc := range testCase {
		t.Log(tc.name)

		result, err := MaxWithLimitation(tc.a, tc.b)
		if tc.isError {
			if err == nil {
				t.Fatal("it should be error")
			}
			continue
		} else {
			if err != nil {
				t.Fatal(err)
			}
		}

		if len(result) != len(tc.expected) {
			t.Fatalf("result is not equal to expected")
		}

		if (result == nil) != (tc.expected == nil) {
			t.Fatalf("result is not equal to expected")
		}

		for i, v := range result {
			if v[0] != tc.expected[i][0] {
				t.Fatalf("result is not equal to expected")
			}
			if v[1] != tc.expected[i][1] {
				t.Fatalf("result is not equal to expected")
			}
		}

	}
}
