package p3956

const inf = 1 << 60

func maximumSum(nums []int, m int, l int, r int) int64 {
	// n <= 1000
	n := len(nums)

	pref := make([]int, n+1)
	for i, v := range nums {
		pref[i+1] = pref[i] + v
	}

	type pair struct {
		first  int
		second int
	}

	dp := make([]int, n+1)

	ans := -inf
	for d := 1; d <= m; d++ {
		var que []pair
		ndp := make([]int, n+1)
		for j := range n + 1 {
			ndp[j] = -inf
		}
		for j := d * l; j <= n; j++ {
			k := j - l
			v := dp[k] - pref[k]
			for len(que) > 0 && last(que).first <= v {
				que = que[:len(que)-1]
			}
			que = append(que, pair{v, k})

			ndp[j] = max(ndp[j-1], que[0].first+pref[j])

			if que[0].second == j-r {
				que = que[1:]
			}
		}
		dp = ndp
		ans = max(ans, dp[n])
	}

	return int64(ans)
}

func last[T any](arr []T) T {
	return arr[len(arr)-1]
}
