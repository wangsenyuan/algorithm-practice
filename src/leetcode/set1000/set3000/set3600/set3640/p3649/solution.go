package p3649

import (
	"sort"
)

func perfectPairs(nums []int) int64 {
	n := len(nums)

	arr := make([]int, n)
	for i := range n {
		arr[i] = abs(nums[i])
	}

	sort.Ints(arr)
	var res int

	for l, r := 0, 0; r < n; r++ {
		for l <= r && arr[l]*2 < arr[r] {
			l++
		}
		res += r - l
	}
	return int64(res)
}

func abs(num int) int {
	return max(num, -num)
}
