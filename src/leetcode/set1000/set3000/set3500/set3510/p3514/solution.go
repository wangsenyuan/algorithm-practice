package p3514

import (
	"math/bits"
	"slices"
)

func uniqueXorTriplets(nums []int) int {
	x := slices.Max(nums)
	h := bits.Len(uint(x))
	occ := make([]bool, 1<<h)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			occ[nums[i]^nums[j]] = true
		}
	}
	var res int
	for x := 0; x < 1<<h; x++ {
		for _, v := range nums {
			if occ[x^v] {
				res++
				break
			}
		}
	}
	return res
}
