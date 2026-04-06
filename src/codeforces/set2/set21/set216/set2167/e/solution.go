package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		_, _, _, res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) (k int, x int, a []int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &k, &x)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	slices.Sort(a)

	res = solve(a, x, k)
	return
}

const inf = 1 << 60

func solve(a []int, x int, k int) []int {

	n := len(a)
	check := func(mid int) bool {
		if mid == 0 {
			return true
		}
		var cnt int
		// 这个k个肯定是放在a[i], a[i+1]的中间
		for i := 0; i < n; i++ {
			if i == 0 {
				// 看看前面可不可以放置
				cnt += max(0, a[i]-mid+1)
			}
			if i+1 < n {
				// i和i+1中间放置
				l := a[i] + mid
				r := a[i+1] - mid
				cnt += max(0, r-l+1)
			} else {
				// i == n - 1
				l := a[i] + mid
				r := x
				cnt += max(0, r-l+1)
			}
		}

		return cnt >= k
	}

	var maxDist int
	for i := 0; i < n; i++ {
		if i == 0 {
			maxDist = a[i]
		}
		if i+1 < n {
			maxDist = max(maxDist, a[i+1]-a[i])
		} else {
			maxDist = max(maxDist, x-a[i])
		}
	}

	lo, hi := 0, maxDist+1
	for lo < hi {
		mid := (lo + hi) / 2
		if !check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	dist := hi - 1
	var res []int
	if dist > 0 {
		add := func(l int, r int) {
			for i := l; i <= r && len(res) < k; i++ {
				res = append(res, i)
			}
		}
		for i := 0; i < n; i++ {
			if i == 0 {
				add(0, a[i]-dist)
			}
			if i+1 < n {
				add(a[i]+dist, a[i+1]-dist)
			} else {
				add(a[i]+dist, x)
			}
		}
	} else {
		for len(res) < k {
			res = append(res, len(res))
		}
	}

	return res[:k]
}
