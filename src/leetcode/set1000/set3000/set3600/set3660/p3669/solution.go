package p3669

import "slices"

func minDifference(n int, k int) []int {
	ans := make([]int, k)
	best := n - 1

	var dfs func(n int, k int)

	var buf []int

	dfs = func(n int, k int) {
		if k == 1 {
			buf = append(buf, n)
			diff := slices.Max(buf) - slices.Min(buf)
			if diff <= best {
				copy(ans, buf)
				best = diff
			}
			buf = buf[:len(buf)-1]
			return
		}
		for x := 1; x*x <= n; x++ {
			if n%x == 0 {
				buf = append(buf, x)
				dfs(n/x, k-1)
				buf = buf[:len(buf)-1]
			}
		}
	}
	dfs(n, k)
	return ans
}
