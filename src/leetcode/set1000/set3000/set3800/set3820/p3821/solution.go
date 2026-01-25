package p3821

func nthSmallest(n int64, k int) int64 {
	C := make([][]int, 51)

	for i := range 51 {
		C[i] = make([]int, 51)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j] + C[i-1][j-1]
		}
	}

	var res int

	u := int(n)
	for i := 50; i >= 0; i-- {
		if C[i][k] >= u {
			continue
		}
		// 这个地方必须是1
		res |= 1 << i
		u -= C[i][k]
		k--
	}

	return int64(res)
}
