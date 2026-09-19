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

func drive(reader *bufio.Reader) float64 {
	var n, v int
	fmt.Fscan(reader, &n, &v)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(v, a, b)
}

func solve(v int, a []int, b []int) float64 {
	// x = min(b[i] / a[i], V / sum(a..))
	var x float64 = 1 << 30
	var sum int
	for i, u := range a {
		sum += u
		x = min(x, float64(b[i])/float64(u))
	}

	x = min(x, float64(v)/float64(sum))

	return x * float64(sum)
}
