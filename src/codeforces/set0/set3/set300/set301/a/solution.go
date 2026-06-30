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
	a := make([]int, 2*n-1)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, a)
}

func solve(n int, a []int) int {
	var sum, neg int
	best := 1 << 30
	for _, v := range a {
		if v < 0 {
			neg++
			v = -v
		}
		sum += v
		best = min(best, v)
	}
	if n%2 == 0 && neg%2 == 1 {
		sum -= 2 * best
	}
	return sum
}
