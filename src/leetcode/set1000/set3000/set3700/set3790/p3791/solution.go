package p3791

import "slices"


func countBalanced(low int64, high int64) int64 {
	low = max(10, low)
	if low > high {
		return 0
	}

	// dp[lz][eq_low][eq_high] = number
	offset := 100
	dp := make([][][][]int, 3)
	ndp := make([][][][]int, 3)
	for i := range 3 {
		dp[i] = make([][][]int, 2)
		ndp[i] = make([][][]int, 2)
		for j := range 2 {
			dp[i][j] = make([][]int, 2)
			ndp[i][j] = make([][]int, 2)
			for k := range 2 {
				dp[i][j][k] = make([]int, offset * 2 + 1)
				ndp[i][j][k] = make([]int, offset * 2 + 1)
			}
		}
	}

	dp[0][1][1][offset] = 1

	a := getDigits(low)
	b := getDigits(high)

	for len(a) < len(b) {
		a = append(a, 0)
	}

	slices.Reverse(a)
	slices.Reverse(b)
	

	for i := range len(a) {
		for pos := range 3 {
			for eq_low := range 2 {
				for eq_high := range 2 {
					for diff := -offset; diff <= offset; diff++ {
						if dp[pos][eq_low][eq_high][diff + offset] == 0 {
							continue
						}
						for d := range 10 {
							if eq_low == 1 && d < a[i] {
								continue
							}
							if eq_high == 1 && d > b[i] {
								break
							}
							new_diff := diff
							var next_pos int
							if pos == 0 {
								if d == 0 {
									next_pos = 0
								} else {
									next_pos = 1
									new_diff = diff + d
								}
							} else if pos == 1 {
								next_pos = 2
								new_diff = diff - d
							} else {
								// pos == 2
								next_pos = 1
								new_diff = diff + d
							}

							neq_low := 1
							if eq_low == 0 || d > a[i] {
								neq_low = 0
							}
							neq_high := 0
							if eq_high == 1 && d == b[i] {
								neq_high = 1
							}
							if new_diff >= -offset && new_diff <= offset {
								ndp[next_pos][neq_low][neq_high][new_diff+offset] += dp[pos][eq_low][eq_high][diff+offset]
							}
						}
					}
				}
			}
		}

		for pos := range 3 {
			for eq_low := range 2 {
				for eq_high := range 2 {
					for diff := -offset; diff <= offset; diff++ {
						dp[pos][eq_low][eq_high][diff+offset] = ndp[pos][eq_low][eq_high][diff+offset]
						ndp[pos][eq_low][eq_high][diff+offset] = 0
					}
				}
			}
		}
	}

	var res int

	for pos := 1; pos <= 2; pos++ {
		for eq_low := range 2 {
			for eq_high := range 2 {
				res += dp[pos][eq_low][eq_high][offset]
			}
		}
	}
	
	return int64(res)
}

func getDigits(num int64) []int {
	var res []int
	for num > 0 {
		res = append(res, int(num%10))
		num /= 10
	}
	return res
}