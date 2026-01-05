package p3795

const mod = 1000000007

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

func numberOfRoutes(grid []string, d int) int {
	n := len(grid)
	m := len(grid[0])

	dp := make([]int, m)
	for i := range m {
		if grid[n-1][i] == '.' {
			dp[i] = 1
		}
	}

	diff := make([]int, m+1)

	checkDist := func(r1 int, c1 int, r2 int, c2 int) bool {
		return (r2-r1)*(r2-r1)+(c2-c1)*(c2-c1) <= d*d
	}

	for i := n - 1; i >= 0; i-- {
		var l, r int
		clear(diff)

		for j := range m {
			for r < m && r-j <= d {
				r++
			}
			diff[j+1] = add(diff[j+1], dp[j])
			diff[r] = sub(diff[r], dp[j])
			for l < j && j-l > d {
				l++
			}
			diff[l] = add(diff[l], dp[j])
			diff[j] = sub(diff[j], dp[j])
		}

		for j := range m {
			diff[j+1] = add(diff[j+1], diff[j])
			if grid[i][j] == '#' {
				diff[j] = 0
			}
		}

		for j := range m {
			diff[j] = add(diff[j], dp[j])
			dp[j] = diff[j]
			if j > 0 {
				diff[j] = add(diff[j], diff[j-1])
			}
		}

		if i > 0 {
			l, r = 0, 0
			for j := range m {
				for r < m && checkDist(i, r, i-1, j) {
					r++
				}
				for l < j && !checkDist(i, l, i-1, j) {
					l++
				}

				if grid[i-1][j] == '.' {

					dp[j] = diff[r-1]
					if l > 0 {
						dp[j] = sub(dp[j], diff[l-1])
					}
				} else {
					dp[j] = 0
				}
			}
		}
	}

	var res int

	for _, v := range dp {
		res = add(res, v)
	}

	return res
}
