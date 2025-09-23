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
	fmt.Println(res[0], res[1])
}

func solve(m int) []int {
	// 先找出124范围内的最小值

	var best int
	var x int

	var f func(m int, steps int, sum int)
	f = func(m int, steps int, sum int) {
		if m == 0 {
			if steps > best || steps == best && sum > x {
				best = steps
				x = sum
			}
			return
		}
		a := sort.Search(min(m+1, 1e5), func(a int) bool {
			return a*a*a > m
		})
		a--
		a3 := a * a * a
		f(m-a3, steps+1, sum+a3)
		if a > 0 {
			a--
			f(a3-1-a*a*a, steps+1, sum+a*a*a)
		}
	}
	f(m, 0, 0)
	return []int{best, x}
}
