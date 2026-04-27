package p3909

func compareBitonicSums(nums []int) int {
	peek := 0
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] && (i == n-1 || nums[i] > nums[i+1]) {
			peek = i
			break
		}
	}
	var sum1 int
	for i := 0; i <= peek; i++ {
		sum1 += nums[i]
	}
	var sum2 int
	for i := peek; i < n; i++ {
		sum2 += nums[i]
	}
	if sum1 > sum2 {
		return 0
	}
	if sum1 < sum2 {
		return 1
	}

	return -1
}
