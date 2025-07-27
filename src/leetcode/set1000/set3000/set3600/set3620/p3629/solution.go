package p3629

import "slices"

func minJumps(nums []int) int {
	var primes []int
	x := slices.Max(nums)

	lpf := make([]int, x+1)
	for i := 2; i <= x; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j > x {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}
	n := len(nums)

	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}

	var que []int
	dp[0] = 0
	que = append(que, 0)

	pos := make([][]int, x+1)

	for i := range nums {
		for y := nums[i]; y > 1; y = y / lpf[y] {
			pos[lpf[y]] = append(pos[lpf[y]], i)
		}
	}

	for len(que) > 0 {
		i := que[0]
		que = que[1:]

		if i > 0 && dp[i-1] == -1 {
			dp[i-1] = dp[i] + 1
			que = append(que, i-1)
		}
		if i+1 < n && dp[i+1] == -1 {
			dp[i+1] = dp[i] + 1
			que = append(que, i+1)
		}
		if lpf[nums[i]] == nums[i] {
			for j := nums[i]; j <= x; j += nums[i] {
				for _, k := range pos[j] {
					if dp[k] == -1 {
						dp[k] = dp[i] + 1
						que = append(que, k)
					}
				}
				pos[j] = pos[j][:0]
			}
		}
	}

	return dp[n-1]
}
