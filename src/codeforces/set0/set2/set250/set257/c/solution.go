package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%.10f\n", drive(reader))
}

func drive(reader *bufio.Reader) float64 {
	var n int
	fmt.Fscan(reader, &n)
	pts := make([][]int, n)
	for i := range n {
		pts[i] = make([]int, 2)
		fmt.Fscan(reader, &pts[i][0], &pts[i][1])
	}
	return solve(pts)
}

func solve(pts [][]int) float64 {
	n := len(pts)
	if n == 1 {
		return 0
	}

	const twoPi = 2 * math.Pi
	angs := make([]float64, n)
	for i, p := range pts {
		a := math.Atan2(float64(p[1]), float64(p[0]))
		if a < 0 {
			a += twoPi
		}
		angs[i] = a
	}
	slices.Sort(angs)

	maxGap := angs[0] + twoPi - angs[n-1]
	for i := 1; i < n; i++ {
		maxGap = max(maxGap, angs[i]-angs[i-1])
	}
	return (twoPi - maxGap) * 180 / math.Pi
}
