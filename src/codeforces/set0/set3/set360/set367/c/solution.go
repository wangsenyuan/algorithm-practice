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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	coupons := make([][]int, m)
	for i := range m {
		coupons[i] = make([]int, 2)
		fmt.Fscan(reader, &coupons[i][0], &coupons[i][1])
	}
	return solve(n, coupons)
}

func solve(n int, coupons [][]int) int {
	m := len(coupons)
	slices.SortFunc(coupons, func(a, b []int) int {
		return b[1] - a[1]
	})

	count := func(x int) int {
		cnt := x * (x - 1) / 2
		if x&1 == 0 {
			// 偶数个点，那么就是奇数条边，需要在每两个点中间，增加一条边，使的所有的点的deg为偶数
			// 但是可以保留两个奇数度的点
			cnt += x/2 - 1
		}
		return cnt
	}

	l, r := 0, m+1
	for l < r {
		mid := (l + r) / 2
		if count(mid) >= n {
			r = mid
		} else {
			l = mid + 1
		}
	}
	r = min(n, r-1)
	var ans int
	for i := range r {
		ans += coupons[i][1]
	}
	return ans
}
