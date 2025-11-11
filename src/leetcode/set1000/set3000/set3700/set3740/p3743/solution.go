package p3743

import "slices"

func maximumScore(nums []int, k int) int64 {
	x := slices.Min(nums)

	for i, v := range nums {
		if v == x {
			if i > 0 {
				shift(nums, i)
			}
			break
		}
	}
	// 第一个是最小值
	play := func(arr []int) int {
		// buy sell
		f := make([][]int, 3)
		for i := range 3 {
			f[i] = make([]int, k+2)
		}
		for i := 1; i <= k+1; i++ {
			f[1][i] = -inf
			f[2][i] = -inf
		}
		f[0][0] = -inf
		for _, v := range arr {
			for j := k + 1; j > 0; j-- {
				f[0][j] = max(f[0][j], f[1][j]+v, f[2][j]-v)
				// 买入
				f[1][j] = max(f[1][j], f[0][j-1]-v)
				// 卖出
				f[2][j] = max(f[2][j], f[0][j-1]+v)
			}
		}
		return f[0][k+1]
	}

	res := play(nums)
	shift(nums, 1)

	res = max(res, play(nums))

	return int64(res)
}

const inf = 1 << 30

func shift(arr []int, k int) {
	reverse(arr[:k])
	reverse(arr[k:])
	reverse(arr)
}

func reverse(arr []int) {
	slices.Reverse(arr)
}
