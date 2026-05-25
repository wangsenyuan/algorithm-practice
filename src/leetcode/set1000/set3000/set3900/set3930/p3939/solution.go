package p3939

const mod = 1_000_000_007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func countValidSubsets(parent []int, nums []int, k int) int {
	n := len(parent)
	adj := make([][]int, n)

	for i := 1; i < n; i++ {
		adj[parent[i]] = append(adj[parent[i]], i)
	}

	fp1 := make([]int, k)
	dp1 := make([]int, k)

	var dfs func(u int) [][]int

	dfs = func(u int) [][]int {
		// 包含u的结果
		dp := make([]int, k)
		dp[nums[u]%k] = 1
		// 不包含u得结果
		fp := make([]int, k)
		fp[0] = 1

		for _, v := range adj[u] {
			sub := dfs(v)

			clear(dp1)
			// 选择u
			for x1, w1 := range sub[1] {
				if w1 != 0 {
					for x2, w2 := range dp {
						x := (x1 + x2) % k
						dp1[x] = add(dp1[x], mul(w1, w2))
					}
				}
			}
			copy(dp, dp1)

			// 不选择u
			clear(fp1)
			for x1 := range k {
				w1 := add(sub[0][x1], sub[1][x1])
				if w1 != 0 {
					for x2 := range k {
						x := (x1 + x2) % k
						fp1[x] = add(fp1[x], mul(w1, fp[x2]))
					}
				}

			}
			copy(fp, fp1)
		}

		return [][]int{dp, fp}
	}
	res := dfs(0)

	ans := add(res[0][0], res[1][0])

	// 减去空集 0 % k == 0
	return sub(ans, 1)
}
