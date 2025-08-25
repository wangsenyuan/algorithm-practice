package p3659

import (
	"slices"
	"sort"
)

func partitionArray(nums []int, k int) bool {
	n := len(nums)
	if n%k != 0 {
		return false
	}
	// 那么分组 = n / k
	sort.Ints(nums)
	// 如果不同的数，没有k个，那么答案是false
	// 每个分组中的数，要不同
	var arr []int
	for i := 0; i < n; {
		j := i
		for i < n && nums[i] == nums[j] {
			i++
		}
		arr = append(arr, i-j)
	}
	x := slices.Max(arr)
	return x <= n/k
}
