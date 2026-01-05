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
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const inf = 1 << 60

func solve(a []int) int {
	n := len(a)

	for i := 1; i < n; i++ {
		a[i] += a[i-1]
	}

	best := a[n-1]
	for i := n - 2; i > 0; i-- {
		best = max(best, a[i]-best)
	}
	return best
}
