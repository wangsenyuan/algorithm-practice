package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.12f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, w int
	fmt.Fscan(reader, &n, &w)
	a := make([]int, 2*n)
	for i := range 2 * n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(w, a)
}

func solve(w int, a []int) float64 {
	n := len(a) / 2
	slices.Sort(a)
	// x = min(a[0], a[n] / 2)
	// and x * 3 * n <= w
	x := float64(w) / float64(3*n)
	x = min(x, float64(a[0]), float64(a[n])/2)
	return x * 3 * float64(n)
}
