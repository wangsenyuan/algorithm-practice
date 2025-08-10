package p3644

func sortPermutation(nums []int) int {
	res := -1
	for i := range len(nums) {
		if nums[i] != i {
			res &= nums[i]
		}
	}
	if res < 0 {
		return 0
	}
	return res
}
