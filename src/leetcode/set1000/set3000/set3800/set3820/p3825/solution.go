package p3825

import "sort"

func longestSubsequence(nums []int) int {
	var res int

	n := len(nums)
	arr := make([]int, n)

	for d := range 30 {
		var p int
		for i := range n {
			if (nums[i]>>d)&1 == 1 {
				j := sort.Search(p, func(j int) bool {
					return arr[j] >= nums[i]
				})
				arr[j] = nums[i]
				if j == p {
					p++
				}
			}
		}

		res = max(res, p)
	}
	return res
}
