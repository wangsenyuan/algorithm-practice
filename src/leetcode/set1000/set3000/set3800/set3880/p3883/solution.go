package p3883

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func countArrays(digitSum []int) int {
	// 9 * 3 + 4 = 31
	ds := make([][]int, 32)
	for _, v := range digitSum {
		if v >= len(ds) {
			return 0
		}
	}

	for i := 0; i <= 5000; i++ {
		s := sumDigits(i)
		ds[s] = append(ds[s], i)
	}

	var dp []int
	for range ds[digitSum[0]] {
		dp = append(dp, 1)
	}
	for i := 1; i < len(digitSum); i++ {
		ws := ds[digitSum[i-1]]
		vs := ds[digitSum[i]]
		ndp := make([]int, len(vs))
		var sum int
		for j1, j2 := 0, 0; j1 < len(vs); j1++ {
			for j2 < len(ws) && ws[j2] <= vs[j1] {
				sum = add(sum, dp[j2])
				j2++
			}
			ndp[j1] = sum
		}
		dp = ndp
	}
	var res int
	for _, v := range dp {
		res = add(res, v)
	}
	return res
}

func sumDigits(x int) int {
	var res int
	for x > 0 {
		res += x % 10
		x /= 10
	}
	return res
}
