package p3876

import "slices"

func uniformArray(nums1 []int) bool {
	n := len(nums1)
	if n == 1 {
		return true
	}
	// 要求nums2[i] 要么都是偶数，要么都是奇数
	// 假设要求是奇数，如果nums1[i]是奇数，使用操作1，没有问题
	// 如果nums1[i]是偶数，那么只要找到一个奇数nums1[j]就可以, 且 nums1[i] >= nums1[j]
	// 所以，如果存在一个奇数，那么true
	// 如果不存在一个奇数，那么全部是偶数
	x := slices.Min(nums1)
	if x&1 == 1 {
		return true
	}
	// x是偶数，那么没法得到奇数的结果，要么全部是偶数
	for _, v := range nums1 {
		if v&1 == 1 {
			// 最小的那个奇数，没法变成偶数
			return false
		}
	}
	return true
}
