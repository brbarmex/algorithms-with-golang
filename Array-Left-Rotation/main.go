package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, d int
	fmt.Scanln(&n, &d)

	var a = make([]string, n, n)
	var pos = n - d
	for i := 0; i < n; i++ {
		fmt.Scan(&a[pos])
		pos = (pos + 1) % n
	}

	fmt.Println(strings.Join(a, " "))
}