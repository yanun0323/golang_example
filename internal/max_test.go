package internal

import (
	"testing"
)

func TestMaxWithLimit(t *testing.T) {
	testCase := []struct {
		name     string
		a, b     []int
		expected []int
		isError  bool
	}{
		{
			name:     "a 比較大",
			a:        []int{5, 0},
			b:        []int{4, 1},
			expected: []int{5, 0},
		},
		{
			name:     "b 比較大",
			a:        []int{5, 3},
			b:        []int{6, 1},
			expected: []int{6, 1},
		},
		{
			name:     "a 比較大(花色)",
			a:        []int{6, 3},
			b:        []int{6, 1},
			expected: []int{6, 3},
		},
		{
			name:     "b 比較大(花色)",
			a:        []int{6, 0},
			b:        []int{6, 1},
			expected: []int{6, 1},
		},
		{
			name:     "a 數字超過",
			a:        []int{14, 1},
			b:        []int{6, 1},
			expected: nil,
			isError:  true,
		},
		{
			name:     "b 數字超過",
			a:        []int{1, 1},
			b:        []int{14, 1},
			expected: nil,
			isError:  true,
		},
		{
			name:     "a <= 0",
			a:        []int{0, 2},
			b:        []int{3, 1},
			expected: nil,
			isError:  true,
		},
		{
			name:     "b <= 0",
			a:        []int{10, 2},
			b:        []int{-1, 1},
			expected: nil,
			isError:  true,
		},
		{
			name:     "a 花色<0",
			a:        []int{1, -2},
			b:        []int{3, 1},
			expected: nil,
			isError:  true,
		},
		{
			name:     "b 花色<0",
			a:        []int{1, 2},
			b:        []int{3, -1},
			expected: nil,
			isError:  true,
		},
		{
			name:     "a 花色超過",
			a:        []int{1, 4},
			b:        []int{3, 1},
			expected: nil,
			isError:  true,
		},
		{
			name:     "b 花色超過",
			a:        []int{1, 2},
			b:        []int{3, 5},
			expected: nil,
			isError:  true,
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
			if v != tc.expected[i] {
				t.Fatalf("result is not equal to expected")
			}
		}

	}
}
