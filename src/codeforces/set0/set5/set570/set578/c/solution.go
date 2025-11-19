package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.10f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const inf = 10000

func solve(a []int) float64 {

	n := len(a)

	get := func(x float64) float64 {
		var sum1, sum2 float64
		var best1 float64 = -inf
		var best2 float64 = inf

		for i := range n {
			u := float64(a[i]) - x
			sum1 += u
			sum2 += u
			best1 = max(best1, sum1)
			best2 = min(best2, sum2)
			if sum1 < 0 {
				sum1 = 0
			}
			if sum2 > 0 {
				sum2 = 0
			}
		}
		return max(best1, -best2)
	}

	var l float64 = -inf
	var r float64 = inf

	for range 100 {
		dist := (r - l) / 3
		m1 := l + dist
		ans1 := get(m1)
		m2 := r - dist
		ans2 := get(m2)
		if ans1 < ans2 {
			r = m2
		} else {
			l = m1
		}
	}
	return get((l + r) / 2)
}
