package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	fmt.Println(res[0], res[1])
}

func drive(reader *bufio.Reader) (a []int, k int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &k)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return a, k, solve(a, k)
}

func solve(a []int, k int) []int {
	n := len(a)
	x := slices.Max(a)

	freq := make([]int, x+1)
	var cnt int

	add := func(x int) {
		freq[x]++
		if freq[x] == 1 {
			cnt++
		}
	}

	rem := func(x int) {
		if freq[x] == 1 {
			cnt--
		}
		freq[x]--
	}

	for l, r := 0, 0; r < n; r++ {
		add(a[r])
		for cnt >= k {
			rem(a[l])
			l++
			if cnt < k {
				add(a[l-1])
				l--
				return []int{l + 1, r + 1}
			}
		}
	}

	return []int{-1, -1}
}
