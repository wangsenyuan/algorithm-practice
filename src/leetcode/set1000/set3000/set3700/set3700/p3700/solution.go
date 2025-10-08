package p3700

func alternatingSum(nums []int) int {
	var res int
	for i, v := range nums {
		if i&1 == 0 {
			res += v
		} else {
			res -= v
		}
	}
	return res
}
