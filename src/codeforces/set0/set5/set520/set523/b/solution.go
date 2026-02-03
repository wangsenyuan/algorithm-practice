package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, cur := range res {
		fmt.Fprintf(writer, "%.6f %.6f %.6f\n", cur[0], cur[1], cur[2])
	}
}

func drive(reader *bufio.Reader) [][]float64 {
	var n, T int
	var c float64
	fmt.Fscan(reader, &n, &T, &c)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	p := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &p[i])
	}
	return solve(T, c, a, p)
}

func solve(T int, c float64, a []int, p []int) [][]float64 {
	n := len(a)
	var ans [][]float64

	var sum float64
	var approx float64
	for i := range n {
		v := float64(a[i])
		sum += v
		if i-T >= 0 {
			sum -= float64(a[i-T])
		}
		real := sum / float64(T)
		approx = (approx + v/float64(T)) / c
		diff := math.Abs(real-approx) / real

		if len(p) > 0 && p[0] == i+1 {
			ans = append(ans, []float64{real, approx, diff})
			p = p[1:]
		}
	}

	return ans
}
