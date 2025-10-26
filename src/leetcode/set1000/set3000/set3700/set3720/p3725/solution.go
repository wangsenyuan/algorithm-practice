package p3725

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func countCoprime(mat [][]int) int {
	X := 151
	dp := make([]int, X)
	ndp := make([]int, X)
	dp[0] = 1
	for _, row := range mat {
		clear(ndp)
		for _, v := range row {
			for x := range X {
				y := gcd(x, v)
				ndp[y] = add(ndp[y], dp[x])
			}
		}
		copy(dp, ndp)
	}
	return dp[1]
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
