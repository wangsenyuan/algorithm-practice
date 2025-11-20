package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var m int
	fmt.Fscan(reader, &m)
	res := solve(m)
	fmt.Println(res)
}

func solve(m int) int {

	play := func(n int) int {
		k := 2
		var res int
		for k*k*k <= n {
			res += n / (k * k * k)
			k++
		}
		return res
	}

	n := sort.Search(1e18, func(n int) bool {
		return play(n) >= m
	})

	if play(n) == m {
		return n
	}
	return -1
}
