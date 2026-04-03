package main

import (
	"fmt"
	"slices"
)

func main() {
	var n, k, m int
	fmt.Scan(&n, &k, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Scan(&a[i])
	}
	res := solve(a, m, k)
	fmt.Printf("%.10f\n", res)
}

func solve(a []int, m int, k int) float64 {
	n := len(a)
	var sum int
	for _, v := range a {
		sum += v
	}

	slices.Sort(a)

	best := float64(sum) / float64(n)

	for rem := 0; rem < m && rem+1 < n; rem++ {
		// 剩余 n - rem - 1 个人可以进行最多 (n - rem - 1) * k
		w := min(m-rem-1, (n-rem-1)*k)
		sum -= a[rem]
		tmp := sum + w
		best = max(best, float64(tmp)/(float64(n-rem-1)))
	}

	return best
}
