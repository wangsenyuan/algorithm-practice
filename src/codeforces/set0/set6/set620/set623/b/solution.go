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
	var n, a, b int
	fmt.Fscan(reader, &n, &a, &b)
	nums := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &nums[i])
	}
	return solve(nums, a, b)
}

const inf = 1 << 60

func solve(nums []int, a int, b int) int {
	// 先得到所有可以做为最小g的质数

	getPrimeFactors := func(x int) []int {
		if x <= 1 {
			return nil
		}
		var res []int
		for i := 2; i <= x; i++ {
			if x%i == 0 {
				res = append(res, i)
				for x%i == 0 {
					x /= i
				}
			}
		}
		if x > 1 {
			res = append(res, x)
		}
		return res
	}

	n := len(nums)

	play := func(x int) int {
		pfs := getPrimeFactors(x - 1)
		pfs = append(pfs, getPrimeFactors(x)...)
		pfs = append(pfs, getPrimeFactors(x+1)...)

		slices.Sort(pfs)
		pfs = slices.Compact(pfs)
		m := len(pfs)
		dp := make([]int, m)
		// fp[i][j] 表示在某删除[l...i]后的最优解
		// fp[i][j] = dp[l][j] + a * (i - l)
		// fp[i][j] = dp[l][j] + a * i - a * l
		//     dp[l][j] - a * l 最小的地方
		fp := make([][]int, n)
		best := make([]int, m)

		for i := range n {
			fp[i] = make([]int, m)

			for j := range m {
				// 全部删除
				fp[i][j] = a*(i+1) + best[j]
				if dp[j] < inf {
					w := pfs[j]
					if (nums[i]-1)%w == 0 || (nums[i]+1)%w == 0 {
						// 且只需要修改
						dp[j] += b
					} else if nums[i]%w != 0 {
						// 没法在前缀得到w的倍数
						dp[j] = inf
					}
				}
				if dp[j] < inf {
					best[j] = min(best[j], dp[j]-a*(i+1))
				}
			}
		}

		ndp := make([]int, m)

		ans := slices.Min(dp)

		for i := n - 1; i >= 0; i-- {
			for j := range m {
				ans = min(ans, fp[i][j]+ndp[j])
				if ndp[j] < inf {
					w := pfs[j]
					if (nums[i]-1)%w == 0 || (nums[i]+1)%w == 0 {
						ndp[j] += b
					} else if nums[i]%w != 0 {
						ndp[j] = inf
					}
				}
			}
		}

		return ans
	}

	u := play(nums[0])
	v := play(nums[n-1])

	return min(u, v)
}
