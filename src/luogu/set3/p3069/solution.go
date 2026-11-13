package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(k, b)
}

func solve(k int, b []int) int {
	freq := make(map[int]int)

	add := func(x int) {
		freq[x]++
	}

	rem := func(x int) {
		freq[x]--
		if freq[x] == 0 {
			delete(freq, x)
		}
	}

	var best int

	for l, r := 0, 0; r < len(b); r++ {
		add(b[r])
		for len(freq) > k+1 {
			rem(b[l])
			l++
		}
		best = max(best, freq[b[r]])
	}

	return best
}
