package p3886

func sortableIntegers(nums []int) int {
	n := len(nums)
	// 3 4 5 1 2 是可以的， 但是 3 4 5 2 1 是不行的

	next := make([]int, n)
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		next[i] = i + 1
		dp[i] = nums[i]
		if i+1 < n && nums[i] <= nums[i+1] {
			next[i] = next[i+1]
		}
		if i+1 < n {
			dp[i] = min(dp[i], dp[i+1])
		}
	}
	fp := make([]int, n)
	for i, v := range nums {
		fp[i] = v
		if i > 0 {
			fp[i] = max(fp[i], fp[i-1])
		}
	}

	check := func(k int) bool {
		for i := n - k; i >= 0; i -= k {
			j := next[i]
			if j < i+k {
				if nums[j] > nums[i] {
					return false
				}
				// nums[j] <= nums[i]
				j1 := next[j]
				if j1 < i+k || nums[i+k-1] > nums[i] {
					return false
				}
			}
			if i > 0 && fp[i-1] > dp[i] {
				return false
			}
		}
		return true
	}

	var res int
	for k := 1; k <= n/k; k++ {
		if n%k == 0 {
			if check(k) {
				res += k
			}
			if k*k != n && check(n/k) {
				res += n / k
			}
		}
	}

	return res
}
