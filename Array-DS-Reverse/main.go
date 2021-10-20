package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"github.com/BrunoBMelo/algorithmArrayReverse/pkg/arrays"
)

func main(){

	var len int
	fmt.Scan(&len)

	arr := make([]int32, len)
	for i := 0; i < len; i++ {
		fmt.Scan(&arr[i])
	}

    res := arrays.Reverse(arr)
	fmt.Println(res)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}