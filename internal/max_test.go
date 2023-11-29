package internal

import "testing"

func TestMaxWithLimit(t *testing.T) {
	testCase := []struct {
		name     string
		a, b     int
		expected int
		isError  bool
	}{
		{
			name:     "a 比較大",
			a:        5,
			b:        4,
			expected: 5,
		},
		{
			name:     "b 比較大",
			a:        5,
			b:        6,
			expected: 6,
		},
		{
			name:     "a 超過",
			a:        55,
			b:        6,
			expected: -1,
			isError:  true,
		},
		{
			name:     "b 超過",
			a:        6,
			b:        55,
			expected: -1,
			isError:  true,
		},
		{
			name:     "a < 0",
			a:        -1,
			b:        2,
			expected: -1,
			isError:  true,
		},
		{
			name:     "b <0",
			a:        2,
			b:        -1,
			expected: -1,
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

		if result != tc.expected {
			t.Fatalf("result is not %d", tc.expected)
		}
	}
}
