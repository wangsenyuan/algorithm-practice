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
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, x, y int
	fmt.Fscan(reader, &n, &x, &y)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, x, y)
}

func solve(a []int, x int, y int) int {
	mx := slices.Max(a)

	n := len(a)
	// 删除	n-1, 只剩一个
	best := n * x

	if mx == 1 {
		mx++
	}

	freq := make([]int, mx+1)
	for _, v := range a {
		freq[v]++
	}

	suf := make([]int, mx+2)
	sum := make([]int, mx+2)
	for i := mx; i > 0; i-- {
		suf[i] = suf[i+1] + freq[i]
		sum[i] = sum[i+1] + freq[i]*i
	}

	for i := 2; i <= mx; i++ {
		// 当gcd = i时
		var cost int
		for j := 0; j <= mx; j += i {

			// 假设有个数 w, w > j, w <= j + i, 将它变成i的倍数
			// 那么需要花费 (j + i - w) * y > x
			j1 := j + i
			j2 := min(mx+1, j1)

			if x <= y {
				cost += (suf[j+1] - suf[j2]) * x
				continue
			}
			d := x / y
			// d * y <= x
			w := j1 - d
			if w <= j {
				cost += ((suf[j+1]-suf[j2])*j1 - (sum[j+1] - sum[j2])) * y
				continue
			}
			w = min(w, j2)
			// w > j
			cost += ((suf[w]-suf[j2])*j1 - (sum[w] - sum[j2])) * y
			cost += (suf[j+1] - suf[w]) * x
		}

		best = min(best, cost)
	}

	return best
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
