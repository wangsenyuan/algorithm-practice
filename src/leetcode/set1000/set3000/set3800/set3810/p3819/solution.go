package p3819

import "slices"

func rotateElements(nums []int, k int) []int {
	var arr []int
	for i, v := range nums {
		if v >= 0 {
			arr = append(arr, i)
		}
	}

	if len(arr) == 0 {
		return nums
	}

	tmp := slices.Clone(arr)

	k %= len(arr)
	k = len(arr) - k

	slices.Reverse(arr[:k])
	slices.Reverse(arr[k:])
	slices.Reverse(arr)

	res := slices.Clone(nums)

	for i := range arr {
		res[arr[i]] = nums[tmp[i]]
	}
	return res
}
