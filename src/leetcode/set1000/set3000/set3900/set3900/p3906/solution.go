package p3906

import "slices"

func countGoodIntegersOnPath(l int64, r int64, directions string) int64 {
	xs := toDecimalDigits(l)
	ys := toDecimalDigits(r)

	special := make([][]bool, 4)
	for i := range 4 {
		special[i] = make([]bool, 4)
	}
	special[0][0] = true
	var x, y int
	for _, d := range directions {
		if d == 'D' {
			x++
		} else {
			y++
		}
		special[x][y] = true
	}
	dp := make([][][][]int, 10)
	ndp := make([][][][]int, 10)
	for i := range 10 {
		dp[i] = make([][][]int, 10)
		ndp[i] = make([][][]int, 10)
		for j := range 10 {
			dp[i][j] = make([][]int, 2)
			ndp[i][j] = make([][]int, 2)
			for k := range 2 {
				dp[i][j][k] = make([]int, 2)
				ndp[i][j][k] = make([]int, 2)
			}
		}
	}

	for s := xs[0]; s <= ys[0]; s++ {
		eqLow := 1
		if s > xs[0] {
			eqLow = 0
		}
		eqHi := 1
		if s < ys[0] {
			eqHi = 0
		}
		dp[s][s][eqLow][eqHi] = 1
	}

	for pos := 1; pos <= 15; pos++ {
		for s1 := range 10 {
			for s2 := range 10 {
				for d1 := range 2 {
					for d2 := range 2 {
						if dp[s1][s2][d1][d2] == 0 {
							continue
						}
						// 在这里填入数字w
						for w := range 10 {
							if w < xs[pos] && d1 == 1 {
								continue
							}
							if w > ys[pos] && d2 == 1 {
								continue
							}

							if w < s1 && special[pos/4][pos%4] {
								continue
							}
							nd1 := d1
							if w > xs[pos] {
								nd1 = 0
							}
							nd2 := d2
							if w < ys[pos] {
								nd2 = 0
							}
							ns1 := s1
							if special[pos/4][pos%4] {
								ns1 = w
							}
							ndp[ns1][w][nd1][nd2] += dp[s1][s2][d1][d2]
						}
					}
				}
			}
		}
		for s1 := range 10 {
			for s2 := range 10 {
				for d1 := range 2 {
					for d2 := range 2 {
						dp[s1][s2][d1][d2] = ndp[s1][s2][d1][d2]
						ndp[s1][s2][d1][d2] = 0
					}
				}
			}
		}
	}

	var ans int
	for s1 := range 10 {
		for d1 := range 2 {
			for d2 := range 2 {
				ans += dp[s1][s1][d1][d2]
			}
		}
	}
	return int64(ans)
}

func toDecimalDigits(n int64) []int {
	var res []int
	for i := n; i > 0; i /= 10 {
		res = append(res, int(i%10))
	}
	for len(res) < 16 {
		res = append(res, 0)
	}
	slices.Reverse(res)
	return res
}
