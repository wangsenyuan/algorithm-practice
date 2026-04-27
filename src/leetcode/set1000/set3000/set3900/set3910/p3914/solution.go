package p3914

import "slices"

func minOperations(nums []int) int64 {
	// nums[0] 不能被增加（它增加了，后面的都要增加）
	var add int
	n := len(nums)
	var res int
	arr := slices.Clone(nums)
	for i := 1; i < n; i++ {
		cur := nums[i] + add
		if cur < arr[i-1] {
			add += arr[i-1] - cur
			res += arr[i-1] - cur
		}
		arr[i] += add
	}
	return int64(res)
}
