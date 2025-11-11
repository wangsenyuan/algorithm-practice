package p3736

import "slices"

func minMoves(nums []int) int {
	x := slices.Max(nums)
	var sum int
	for _, v := range nums {
		if v < x {
			sum += x - v
		}
	}
	return sum
}
