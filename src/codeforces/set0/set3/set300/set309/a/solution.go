package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Printf("%.6f\n", res)
}

func process(reader *bufio.Reader) float64 {
	var n, l, t int
	fmt.Fscan(reader, &n, &l, &t)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	return solve(l, t, a)
}

func solve(l int, t int, a []int) float64 {
	n := len(a)
	// l是长度，t是总时间
	k, p := t/l, t%l
	ans := k * n * (n - 1) / 2

	if p == 0 {
		return float64(ans)
	}

	dist := func(i int, j int) int {
		if j < n {
			return a[j] - a[i]
		}
		return a[j%n] + l - a[i]
	}

	var more int
	// x 表示距离,
	for i, j, k := 0, 0, 0; i < n; i++ {
		// i向前跑，前面的往后跑，相对运动
		for j < 2*n && j < i+n && dist(i, j) <= 2*p {
			j++
		}

		more += j - i - 1

		if 2*p > l {
			for k < 2*n && k < i+n && dist(i, k) <= 2*p-l {
				k++
			}

			more += k - i - 1
		}

	}

	res := float64(ans)
	res += float64(more) / 4

	return res
}
