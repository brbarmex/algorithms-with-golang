package hourglass

import "testing"

func TestHourglassSum(t *testing.T) {
	var cases = []struct {
		arr      [][]int
		expected int
	}{
		{
			arr: [][]int{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 0, 2, 4, 4, 0},
				{0, 0, 0, 2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			},
			expected: 19,
		},
		{
			arr: [][]int{
				{0, -4, -6, 0, -7, -6},
				{-1, -2, -6, -8, -3, -1},
				{-8, -4, -2, -8, -8, -6},
				{-3, -1, -2, -5, -7, -4},
				{-3, -5, -3, -6, -6, -6},
				{-3, -6, 0, -8, -6, -7},
			},
			expected: -19,
		},
		{
			arr: [][]int{
				{-1, 1, -1, 0, 0, 0},
				{0, -1, 0, 0, 0, 0},
				{-1, -1, -1, 0, 0, 0},
				{0, -9, 2, -4, -4, 0},
				{-7, 0, 0, -2, 0, 0},
				{0, 0, -1, -2, -4, 0},
			},
			expected: 0,
		},
	}

	for _, options := range cases {
		got := Sum(options.arr)
		if got != options.expected {
			t.Errorf("HourglassSum(arr) = %d; want %d", got, options.expected)
		}
	}
}
