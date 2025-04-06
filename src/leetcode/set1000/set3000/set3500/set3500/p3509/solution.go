package main

func maxProduct(nums []int, k, limit int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	if total < abs(k) {
		return -1
	}

	ans := -1
	type args struct {
		i, s, m    int
		odd, empty bool
	}
	vis := map[args]bool{}
	var dfs func(int, int, int, bool, bool)
	dfs = func(i, s, m int, odd, empty bool) {
		if ans == limit { // 已经达到最大值
			return
		}

		if i == len(nums) {
			if !empty && s == k && m <= limit {
				ans = max(ans, m)
			}
			return
		}

		t := args{i, s, m, odd, empty}
		if vis[t] {
			return
		}
		vis[t] = true

		// 不选 x
		dfs(i+1, s, m, odd, empty)

		// 选 x
		x := nums[i]
		if odd {
			s -= x
		} else {
			s += x
		}
		dfs(i+1, s, min(m*x, limit+1), !odd, false)
	}
	dfs(0, 0, 1, false, true)
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
