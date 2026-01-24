package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res < 0 {
		fmt.Println(-1)
	} else {
		fmt.Printf("%.8f\n", res)
	}
}

func drive(reader *bufio.Reader) float64 {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}

	return solve(m, a, b)
}

func solve(m int, a []int, b []int) float64 {
	M := float64(m)
	n := len(a)
	check := func(x float64) bool {
		// 能否用x的燃料，完成旅行
		for i := range n {
			// 当前是x的，假设使用了其中x0的部分
			// 那么 x0 * a[i] >= m + x
			x0 := (M + x) / float64(a[i])
			if x0 > x {
				return false
			}
			// 到达空中后，剩余的部分
			x -= x0
			j := (i + 1) % n
			// 使用y0的燃料到达j
			// y0 * b[j] >= m + x
			y0 := (M + x) / float64(b[j])
			if y0 > x {
				return false
			}
			x -= y0
		}
		return true
	}

	var lo, hi float64 = 0, 1 << 60
	for range 100 {
		mid := (lo + hi) / 2
		if check(mid) {
			hi = mid
		} else {
			lo = mid
		}
	}

	res := (lo + hi) / 2

	if math.Abs(res-1<<60) < 1e-6 {
		return -1
	}

	return res
}
