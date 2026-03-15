package p3866

import "slices"

func firstUniqueEven(nums []int) int {
	mx := slices.Max(nums)
	freq := make([]int, mx+1)
	for _, num := range nums {
		freq[num]++
	}

	for _, num := range nums {
		if num%2 == 0 && freq[num] == 1 {
			return num
		}
	}
	return -1
}
