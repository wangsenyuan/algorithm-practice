package p3698

func zigZagArrays(n int, l int, r int) int {
	d := r - l + 1
	dp := make([][2]int, d)
	for i := 0; i < d; i++ {
		dp[i][0] = d - i - 1
		dp[i][1] = i
	}

	for range n - 2 {
		ndp := make([][2]int, d)

		var sum int
		for i := range d {
			ndp[i][1] = sum
			sum = modAdd(sum, dp[i][0])
		}
		sum = 0
		for i := d - 1; i >= 0; i-- {
			ndp[i][0] = sum
			sum = modAdd(sum, dp[i][1])
		}

		dp = ndp
	}

	var res int

	for i := range d {
		res = modAdd(res, dp[i][0])
		res = modAdd(res, dp[i][1])
	}

	return res
}

const MOD = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	return a * b % MOD
}
