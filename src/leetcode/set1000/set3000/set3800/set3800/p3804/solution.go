package p3804

func centeredSubarrays(nums []int) int {
	var res int

	n := len(nums)
	freq := make(map[int]int)

	for i := range n {
		var sum int
		clear(freq)
		for j := i; j < n; j++ {
			freq[nums[j]]++
			sum += nums[j]
			if freq[sum] > 0 {
				res++
			}
		}
	}
	return res
}
