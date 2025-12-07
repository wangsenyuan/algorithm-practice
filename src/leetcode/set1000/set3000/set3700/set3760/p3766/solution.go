package p3766

import (
	"slices"
	"sort"
)

func minOperations(nums []int) []int {
	var arr []int
	for i := 1; i < 10240; i += 2 {
		j := reverse(i)
		if i == j {
			arr = append(arr, i)
		}
	}

	slices.Sort(arr)
	ans := make([]int, len(nums))
	for i, num := range nums {
		j := sort.SearchInts(arr, num)
		ans[i] = arr[j] - num
		if j > 0 {
			ans[i] = min(ans[i], num-arr[j-1])
		}
	}
	return ans
}

func reverse(num int) int {
	var res int
	for num > 0 {
		x := num & 1
		res = (res << 1) | x
		num >>= 1
	}
	return res
}
