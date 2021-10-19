package hourglass

import "math"

func Sum(arr [][]int) int {
	a, b, c, d, bigger := -1, 0, 1, 0, math.MinInt32

	for h := 0; h < 4; h++ {
		for v := 0; v < 4; v++ {
			a, b, c, d = a+1, b+1, c+1, d+1
			if x := arr[h][a] + arr[h][b] + arr[h][c] + arr[h+1][d] + arr[h+2][a] + arr[h+2][b] + arr[h+2][c]; x > bigger {
				bigger = x
			}
		}
		a, b, c, d = -1, 0, 1, 0
	}
	return bigger
}
