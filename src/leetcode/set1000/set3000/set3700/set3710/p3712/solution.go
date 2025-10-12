package p3712

import "slices"

func sumDivisibleByK(nums []int, k int) int {
	x := slices.Max(nums)
	freq := make([]int, x+1)
	for _, v := range nums {
		freq[v]++
	}
	var sum int
	for i := 1; i <= x; i++ {
		if freq[i]%k == 0 {
			sum += i * freq[i]
		}
	}
	return sum
}
