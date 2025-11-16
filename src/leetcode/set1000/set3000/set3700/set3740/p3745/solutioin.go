package p3745

func maximizeExpressionOfThree(nums []int) int {
	c := -1
	for i := range nums {
		if c == -1 || nums[i] < nums[c] {
			c = i
		}
	}
	a, b := -1, -1
	for i := range nums {
		if i == c {
			continue
		}
		if a == -1 || nums[i] >= nums[a] {
			b = a
			a = i
		} else if b == -1 || nums[i] >= nums[b] {
			b = i
		}
	}
	return nums[a] + nums[b] - nums[c]
}
