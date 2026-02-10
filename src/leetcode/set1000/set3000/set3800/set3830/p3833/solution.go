package p3833

func dominantIndices(nums []int) int {

	n := len(nums)
	var res int
	var sum int
	for i := n - 1; i >= 0; i-- {
		v := nums[i]
		if (n-i-1)*v > sum {
			res++
		}
		sum += v
	}
	return res
}
