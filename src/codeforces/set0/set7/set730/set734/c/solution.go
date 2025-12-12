package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	var x, s int
	fmt.Fscan(reader, &x, &s)
	a := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	c := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &c[i])
	}
	d := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &d[i])
	}
	return solve(n, x, s, a, b, c, d)
}

type pair struct {
	first  int
	second int
}

func solve(n int, x int, s int, a []int, b []int, c []int, d []int) int {
	best := n * x

	m := len(a)
	arr := make([]pair, m)
	for i := range m {
		arr[i] = pair{a[i], b[i]}
	}
	slices.SortFunc(arr, func(x pair, y pair) int {
		return x.second - y.second
	})

	dp := make([]int, m)
	for i := range m {
		dp[i] = arr[i].first
		if i > 0 {
			dp[i] = min(dp[i], dp[i-1])
		}
		x1 := min(x, dp[i])
		if arr[i].second <= s {
			best = min(best, x1*n)
		}
	}

	k := len(c)

	for i := range k {
		if d[i] > s {
			break
		}
		n1 := n - c[i]
		if n1 <= 0 {
			// 全部立刻完成了
			return 0
		}
		// 还有s1的能量
		s1 := s - d[i]
		j := sort.Search(m, func(j int) bool {
			return arr[j].second > s1
		})
		j--
		if j >= 0 {
			x1 := min(x, dp[j])
			best = min(best, x1*n1)
		} else {
			// 即使找到合适的方式1，也可以用x的时间准备剩余的部分
			best = min(best, x*n1)
		}
	}

	return best
}
