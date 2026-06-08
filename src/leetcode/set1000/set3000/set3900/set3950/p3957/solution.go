package p3957

const inf = 1 << 60

type pair struct {
	sum int64
	cnt int
}

func maximumSum(nums []int, m int, l int, r int) int64 {
	res := calc(nums, l, r, 0)
	if res.cnt == 0 {
		return bestOne(nums, l, r)
	}
	if res.cnt <= m {
		return res.sum
	}

	lo, hi := int64(-inf), int64(inf)
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if calc(nums, l, r, mid).cnt >= m {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	res = calc(nums, l, r, lo)
	return res.sum + lo*int64(m)
}

func calc(nums []int, l int, r int, cost int64) pair {
	n := len(nums)
	pref := make([]int64, n+1)
	for i, num := range nums {
		pref[i+1] = pref[i] + int64(num)
	}

	dp := make([]pair, n+1)
	que := make([]state, n+1)
	var front, tail int

	for i := 1; i <= n; i++ {
		j := i - l
		if j >= 0 {
			cur := state{pair{dp[j].sum - pref[j], dp[j].cnt}, j}
			for front < tail && better(cur.val, que[tail-1].val) {
				tail--
			}
			que[tail] = cur
			tail++
		}

		low := i - r
		for front < tail && que[front].idx < low {
			front++
		}

		dp[i] = dp[i-1]
		if front < tail {
			cur := pair{
				sum: pref[i] + que[front].val.sum - cost,
				cnt: que[front].val.cnt + 1,
			}
			if better(cur, dp[i]) {
				dp[i] = cur
			}
		}
	}

	return dp[n]
}

type state struct {
	val pair
	idx int
}

func better(a pair, b pair) bool {
	return a.sum > b.sum || a.sum == b.sum && a.cnt > b.cnt
}

func bestOne(nums []int, l int, r int) int64 {
	n := len(nums)
	pref := make([]int64, n+1)
	for i, num := range nums {
		pref[i+1] = pref[i] + int64(num)
	}

	ans := int64(-inf)
	que := make([]int, n+1)
	var front, tail int

	for i := 1; i <= n; i++ {
		j := i - l
		if j >= 0 {
			for front < tail && pref[j] <= pref[que[tail-1]] {
				tail--
			}
			que[tail] = j
			tail++
		}

		low := i - r
		for front < tail && que[front] < low {
			front++
		}

		if front < tail {
			ans = max(ans, pref[i]-pref[que[front]])
		}
	}

	return ans
}
