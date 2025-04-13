package p3513

import (
	"math/bits"
)

func uniqueXorTriplets(nums []int) int {
	n := len(nums)
	if n < 3 {
		return n
	}
	h := bits.Len(uint(n))
	return 1 << h
}
