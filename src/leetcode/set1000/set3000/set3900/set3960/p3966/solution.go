package p3966

import "slices"

func goodIntegers(l int64, r int64, k int) int64 {
	res := play(int(r), k)
	if l > 0 {
		res -= play(int(l-1), k)
	}
	return int64(res)
}

func play(num int, k int) int {
	var ds []int
	for i := num; i > 0; i /= 10 {
		ds = append(ds, int(i%10))
	}
	slices.Reverse(ds)
	n := len(ds)
	var dp [10][2]int
	for i := 1; i < ds[0]; i++ {
		dp[i][0] = 1
	}
	dp[ds[0]][1] = 1

	for i := 1; i < n; i++ {
		v := ds[i]
		var ndp [10][2]int
		for j := range 10 {
			for d := range 2 {
				if dp[j][d] == 0 {
					continue
				}
				for j1 := 0; j1 < 10; j1++ {
					if d == 1 && j1 > v {
						break
					}
					if abs(j-j1) <= k {
						d1 := d
						if d == 1 && j1 < v {
							d1 = 0
						}
						ndp[j1][d1] += dp[j][d]
					}
				}
			}
		}
		for j := 1; j < 10; j++ {
			ndp[j][0]++
		}
		dp = ndp
	}
	var res int
	for j := range 10 {
		for d := range 2 {
			res += dp[j][d]
		}
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}
