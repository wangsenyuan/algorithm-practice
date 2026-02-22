package p3850

func countSequences(nums []int, k int64) int {
	cnt := make([]int, 7)
	for _, num := range nums {
		if num == 1 || num == 5 {
			cnt[num]++
			continue
		}
		for num%2 == 0 {
			cnt[2]++
			num /= 2
		}
		for num%3 == 0 {
			cnt[3]++
			num /= 3
		}
	}

	freq := make([]int, 6)
	kk := int(k)
	for i := 2; i <= 5; i++ {
		for kk%i == 0 {
			freq[i]++
			kk /= i
		}
	}
	if kk > 1 {
		// 有除了 2, 3, 5 外的其他的质数因子
		return 0
	}
	for _, i := range []int{2, 3, 5} {
		if freq[i] > cnt[i] {
			return 0
		}
	}
	// 其中x个作为除数
	// 2和3，有点麻烦，
	// dp[i][j] 表示有i个2，j个3的组合数
	dp := make([][]int, 2*cnt[2]+1)
	ndp := make([][]int, 2*cnt[2]+1)
	for i := range dp {
		dp[i] = make([]int, 2*cnt[3]+1)
		ndp[i] = make([]int, 2*cnt[3]+1)
	}
	dp[cnt[2]][cnt[3]] = 1
	for _, num := range nums {
		if num == 1 || num == 5 {
			continue
		}
		var u, v int
		for num%2 == 0 {
			u++
			num /= 2
		}
		for num%3 == 0 {
			v++
			num /= 3
		}

		for c2 := range 2*cnt[2] + 1 {
			for c3 := range 2*cnt[3] + 1 {
				if dp[c2][c3] == 0 {
					continue
				}
				// 如果不选择num
				ndp[c2][c3] += dp[c2][c3]
				// 如果选择乘法
				if c2+u <= 2*cnt[2] && c3+v <= 2*cnt[3] {
					ndp[c2+u][c3+v] += dp[c2][c3]
				}
				// 如果作为除法
				if c2 >= u && c3 >= v {
					ndp[c2-u][c3-v] += dp[c2][c3]
				}
			}
		}

		for c2 := range 2*cnt[2] + 1 {
			for c3 := range 2*cnt[3] + 1 {
				dp[c2][c3] = ndp[c2][c3]
				ndp[c2][c3] = 0
			}
		}
	}

	C := make([][]int, cnt[5]+1)
	for i := range cnt[5] + 1 {
		C[i] = make([]int, cnt[5]+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j] + C[i-1][j-1]
		}
	}

	res1 := dp[freq[2]+cnt[2]][freq[3]+cnt[3]]
	var res2 int
	// 质因子5，需要特殊处理, 比如选择 C(cnt[5], 2 * x + freq[5])
	for x := 0; 2*x+freq[5] <= cnt[5]; x++ {
		res2 += C[cnt[5]][2*x+freq[5]] * C[2*x+freq[5]][x]
	}

	res3 := 1
	for range cnt[1] {
		res3 *= 3
	}

	return res1 * res2 * res3
}
