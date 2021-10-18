package main

import (
    "fmt"
    "github.com/BrunoBMelo/algorithm2DArrayDS/pkg/hourglass"
)

func main() {
	arr := make([][]int32, 6)
	for i := 0; i < 6; i++ {
		arr[i] = make([]int32, 6)
		for j := 0; j < 6; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

    var result = hourglass.Sum(arr)
    fmt.Println(result)
}