package p3687

func evenNumberBitwiseORs(nums []int) int {
	var res int
	for _, v := range nums {
		if v&1 == 0 {
			res |= v
		}
	}
	return res
}
