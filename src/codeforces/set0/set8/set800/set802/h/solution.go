package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	res := solve(n)
	fmt.Println(res[0], res[1])
}

func solve(n int) []string {
	c := make([]int, 200)
	for i := 3; i <= 199; i++ {
		c[i] = i * (i - 1) * (i - 2) / 6
	}

	var arr []int
	for i := 199; i >= 3; i-- {
		for n >= c[i] {
			arr = append(arr, i)
			n -= c[i]
		}
	}

	var s []byte
	var cnt int
	for i := len(arr) - 1; i >= 0; i-- {
		for cnt < arr[i] {
			s = append(s, 'a')
			cnt++
		}
		s = append(s, 'b')
	}

	// 每段任选3个，

	return []string{string(s), "aaab"}
}
