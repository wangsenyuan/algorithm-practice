package p3718

import "slices"

func missingMultiple(nums []int, k int) int {
	x := slices.Max(nums)

	marked := make([]bool, x+1)
	for _, v := range nums {
		marked[v] = true
	}

	for k1 := k; ; k1 += k {
		if k1 > x || !marked[k1] {
			return k1
		}
	}
}
