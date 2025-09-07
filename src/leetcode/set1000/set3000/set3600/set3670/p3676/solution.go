package p3676

import "slices"

func bowlSubarrays(nums []int) int64 {
	res := solve(nums)
	slices.Reverse(nums)
	res += solve(nums)

	return int64(res)
}

func solve(arr []int) int {
	n := len(arr)
	stack := make([]int, n)
	var top int
	var res int
	for i, v := range arr {
		for top > 0 && arr[stack[top-1]] < v {
			top--
		}
		if top > 0 && i-stack[top-1] >= 2 {
			res++
		}
		stack[top] = i
		top++
	}
	return res
}
