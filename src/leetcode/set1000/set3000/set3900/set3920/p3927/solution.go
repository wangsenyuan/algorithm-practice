package p3927

import "slices"

func minArraySum(nums []int) int64 {
	x := slices.Max(nums)
	freq := make([]int, x+1)
	for _, v := range nums {
		freq[v]++
	}
	n := len(nums)
	if freq[1] > 0 {
		return int64(n)
	}
	var res int
	marked := make([]bool, x+1)
	for i := 2; i <= x; i++ {
		if marked[i] || freq[i] == 0 {
			continue
		}
		var cnt int
		for j := i; j <= x; j += i {
			if !marked[j] {
				cnt += freq[j]
			}
			marked[j] = true
		}
		res += cnt * i
	}
	return int64(res)
}
