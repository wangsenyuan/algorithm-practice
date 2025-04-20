package p3522

func maximumPossibleSize(nums []int) int {
	n := len(nums)
	prev := make([]int, n)

	var res int

	for i, v := range nums {
		if i == 0 || v >= nums[prev[i-1]] {
			prev[i] = i
		} else {
			prev[i] = prev[i-1]
		}
		if prev[i] == i {
			res++
		}
	}

	return res
}
