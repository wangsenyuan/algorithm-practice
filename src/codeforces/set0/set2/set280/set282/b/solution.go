package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		var ignore int
		fmt.Fscan(reader, &a[i], &ignore)
	}
	return solve(a)
}

func solve(a []int) string {
	n := len(a)
	buf := make([]byte, n)

	var sum int
	for i, v := range a {
		if sum+v <= 500 {
			buf[i] = 'A'
			sum += v
		} else {
			sum -= 1000 - v
			buf[i] = 'G'
		}
	}
	return string(buf)
}
