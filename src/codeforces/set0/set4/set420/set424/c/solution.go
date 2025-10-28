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

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(p)
}

func solve(p []int) int {
	n := len(p)

	var res int
	diff := make([]int, n+2)

	for i, v := range p {
		res ^= v
		i++
		x, r := n/i, n%i
		if r == 0 {
			if x&1 == 1 {
				diff[1]++
				diff[i]--
			}
		} else {
			// 比如 i = 2, n = 5, x = 2, r = 1
			// [1, 2], [3, 4], [5]
			if x&1 == 1 {
				diff[r+1]++
				diff[i]--
			} else {
				diff[1]++
				diff[r+1]--
			}
		}
	}

	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
		if diff[i]&1 == 1 {
			res ^= i
		}
	}
	return res
}
