package p3912

func findValidElements(nums []int) []int {
	n := len(nums)
	ok := make([]bool, n)
	prev := -1
	for i, v := range nums {
		if v > prev {
			ok[i] = true
		}
		prev = max(prev, v)
	}
	next := -1
	for i := n - 1; i >= 0; i-- {
		if nums[i] > next {
			ok[i] = true
		}
		next = max(next, nums[i])
	}
	var res []int
	for i, v := range nums {
		if ok[i] {
			res = append(res, v)
		}
	}
	return res
}
