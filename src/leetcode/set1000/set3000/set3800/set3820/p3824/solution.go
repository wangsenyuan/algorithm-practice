package p3824

import "sort"

func minimumK(nums []int) int {

	check := func(k int) bool {
		if k == 0 {
			return false
		}
		var sum int
		for _, v := range nums {
			w := (v + k - 1) / k
			sum += w
		}
		return sum <= k*k
	}

	return sort.Search(1<<30, check)
}
