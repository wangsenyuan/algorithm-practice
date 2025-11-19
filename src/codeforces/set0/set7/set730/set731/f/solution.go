package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	// n := len(a)
	x := slices.Max(a)
	freq := make([]int, x+1)
	for _, v := range a {
		freq[v]++
	}
	for v := 1; v <= x; v++ {
		freq[v] += freq[v-1]
	}

	var best int

	for v := 1; v <= x; v++ {
		if freq[v]-freq[v-1] == 0 {
			continue
		}
		var sum int
		for w := v; w <= x; w += v {
			sum += (freq[min(x, w+v-1)] - freq[w-1]) * w
		}
		best = max(best, sum)
	}

	return best
}
