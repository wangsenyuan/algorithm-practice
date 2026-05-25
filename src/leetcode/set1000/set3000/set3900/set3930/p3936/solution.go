package p3936

func minimumSwaps(nums []int) int {
	var cnt int
	for _, v := range nums {
		if v == 0 {
			cnt++
		}
	}
	var res int
	for i := len(nums) - 1; i >= 0 && cnt > 0; i-- {
		if nums[i] != 0 {
			res++
		}
		cnt--
	}

	return res
}
