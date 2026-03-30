package p3880

func minAbsoluteDifference(nums []int) int {
	last := make([]int, 3)
	for i := range 3 {
		last[i] = -1
	}
	res := len(nums)
	for i, v := range nums {
		if v == 2 {
			if last[1] >= 0 {
				res = min(res, i - last[1])
			}
		} else if v == 1 {
			if last[2] >= 0 {
				res = min(res, i - last[2])
			}
		}
		last[v] = i
	}
	if res == len(nums) {
		return -1
	}
	return res
}
