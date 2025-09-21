package p3689

type pair struct {
	first  int
	second int
}

func minSplitMerge(nums1 []int, nums2 []int) int {
	n := len(nums1)

	vis := make(map[string]bool)
	dp := make(map[string]int)

	var dfs func(state string) int

	dfs = func(state string) int {
		// state表示的状态和nums2是否一致
		if v, ok := dp[state]; ok {
			return v
		}
		var diff int
		buf := make([]int, n)
		for i := 0; i < n; i++ {
			buf[i] = int(state[i] - '0')
			if nums1[buf[i]] != nums2[i] {
				diff++
			}
		}
		if diff == 0 {
			return 0
		}
		vis[state] = true
		res := inf
		// 选择一段区间
		for l := 0; l < n; l++ {
			for r := l; r < n; r++ {
				s1 := state[0:l]
				s3 := state[r+1:]
				s2 := state[l : r+1]
				for i := 0; i <= len(s1); i++ {
					next := s1[:i]
					next += s2
					next += s1[i:]
					next += s3
					if !vis[next] {
						res = min(res, dfs(next)+1)
					}
				}
				for i := 0; i <= len(s3); i++ {
					next := s1 + s3[:i]
					next += s2
					next += s3[i:]
					if !vis[next] {
						res = min(res, dfs(next)+1)
					}
				}
			}
		}
		vis[state] = false

		dp[state] = res
		return res
	}

	state := make([]byte, n)
	for i := 0; i < n; i++ {
		state[i] = byte(i + '0')
	}

	return dfs(string(state))
}

const inf = 1 << 30
