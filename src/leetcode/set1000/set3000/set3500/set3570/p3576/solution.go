package p3576

func canMakeEqual(nums []int, k int) bool {
	n := len(nums)
	diff := make([]int, n+1)
	check := func(x int) bool {
		clear(diff)
		var cnt int
		for i := 0; i < n; i++ {
			if i > 0 {
				diff[i] += diff[i-1]
			}
			if diff[i]&1 == 0 && nums[i] != x || diff[i]&1 == 1 && nums[i] == x {
				if i+1 == n {
					// 无法操作
					return false
				}
				// 需要操作
				cnt++
				diff[i]++
				diff[i+2]--
			}
		}
		return cnt <= k
	}
	return check(1) || check(-1)
}
