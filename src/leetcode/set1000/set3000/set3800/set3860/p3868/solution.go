package p3868

import "slices"

func minCost(nums1 []int, nums2 []int) int {
	mx := max(slices.Max(nums1), slices.Max(nums2))
	freq := make([]int, mx+1)
	f1 := make([]int, mx+1)
	for _, num := range nums1 {
		freq[num]++
		f1[num]++
	}
	f2 := make([]int, mx+1)
	for _, num := range nums2 {
		freq[num]++
		f2[num]++
	}
	var ans int

	for _, v := range nums1 {
		if f1[v] == 0 {
			continue
		}
		if freq[v]%2 != 0 {
			return -1
		}
		w := freq[v] / 2
		ans += abs(w - f1[v])
		f1[v] = 0
		f2[v] = 0
	}

	for _, v := range nums2 {
		if f2[v] == 0 {
			continue
		}
		if freq[v]%2 != 0 {
			return -1
		}
		w := freq[v] / 2

		ans += abs(w - f2[v])
		f2[v] = 0
	}

	if ans%2 != 0 {
		return -1
	}

	return ans / 2
}

func abs(num int) int {
	return max(num, -num)
}
