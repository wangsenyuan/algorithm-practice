package p3566

func checkEqualPartitions(nums []int, target int64) bool {
	n := len(nums)

	N := 1 << n

	for state := 1; state < N; state++ {
		p1, p2 := 1, 1
		for i := range n {
			if (state>>i)&1 == 1 {
				p1 *= nums[i]
			} else {
				p2 *= nums[i]
			}
			if int64(p1) > target || int64(p2) > target {
				break
			}
		}
		if p1 == p2 && int64(p1) == target {
			return true
		}
	}
	return false
}
