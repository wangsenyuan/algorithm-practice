package p3621

import "math/bits"

func popcountDepth(n int64, k int) int64 {
	if k == 0 {
		return 1
	}

	f := make([]int, 61)

	f[1] = 0

	var arr []int
	for i := 2; i <= 60; i++ {
		j := bits.OnesCount(uint(i))
		f[i] = f[j] + 1
		if f[i] == k-1 {
			arr = append(arr, i)
		}
	}
	if k == 1 {
		arr = append(arr, 1)
	}

	h := bits.Len(uint(n))

	C := make([][]int, h)
	for i := range h {
		C[i] = make([]int, i+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}

	var cnt int

	var res int
	for i := h - 1; i >= 0; i-- {
		x := (n >> i) & 1
		if x == 1 {
			// 如果这里放置0, 然后在后i位，放置 w - cnt 个 1
			for _, w := range arr {
				// cnt + y = w
				if cnt <= w && w-cnt <= i {
					res += C[i][w-cnt]
				}
				// 将1舍弃掉
				if w == 1 && cnt == 0 {
					res--
				}
			}

			cnt++
		}
	}
	if f[cnt] == k-1 {
		res++
	}
	return int64(res)
}
