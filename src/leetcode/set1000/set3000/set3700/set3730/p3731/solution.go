package p3731

import "sort"

func findMissingElements(nums []int) []int {
	sort.Ints(nums)
	var res []int
	n := len(nums)
	var j int
	for i := nums[0]; i < nums[n-1]; i++ {
		if j < n && nums[j] == i {
			j++
			continue
		}
		res = append(res, i)
	}
	return res
}
