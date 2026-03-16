package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var k, b, n, t int
	fmt.Fscan(reader, &k, &b, &n, &t)
	return solve(k, b, n, t)
}

func solve(k int, b int, n int, t int) int {
	if t == 1 {
		return n
	}
	// t >= 1
	check := func(x int) bool {
		// 经过时刻x, 是否有t个
		cur := 1
		for range x {
			cur = cur*k + b
			if cur > t {
				return true
			}
		}
		return false
	}

	if !check(n) {
		return 0
	}

	n1 := sort.Search(n, check) - 1
	return n - n1
}
