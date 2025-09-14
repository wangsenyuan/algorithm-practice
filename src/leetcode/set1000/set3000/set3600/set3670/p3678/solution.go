package p3678

import "sort"

func smallestAbsent(nums []int) int {
	sort.Ints(nums)
	var sum int
	for _, num := range nums {
		sum += num
	}
	n := len(nums)
	x := 1
	if sum > 0 {
		x = (sum + n - 1) / n
		if x*n == sum {
			x++
		}
	}
	for {
		i := sort.SearchInts(nums, x)
		if i == len(nums) || nums[i] > x {
			return x
		}
		x++
	}
}
