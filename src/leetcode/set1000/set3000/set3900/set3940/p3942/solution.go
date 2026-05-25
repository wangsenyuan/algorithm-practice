package p3942

const inf = 1 << 60

func minOperations(nums []int) int {
	n := len(nums)
	// 0 1, 2, ... n-1 升序转换而来,
	// 或者是n-1, ... 0 降序转换而来
	var cnt int
	var p int
	for i := 1; i < n; i++ {
		if nums[i-1] > nums[i] {
			cnt++
			p = i
		}
	}

	if cnt == 0 {
		return 0
	}
	ans := inf
	if cnt == 1 && nums[0] > nums[n-1] {
		// 3 4 5 0 1 2
		// 使用p次，将0移动到头部
		// 或者先翻转一次，使用n-p次将n-1移动到头部，再翻转一次
		ans = min(p, n-p+2)
	}
	p = 0
	cnt = 0
	for i := 1; i < n && cnt < 2; i++ {
		if nums[i-1] < nums[i] {
			cnt++
			p = i
		}
	}
	if cnt == 0 {
		return 1
	}
	if cnt == 1 && nums[0] < nums[n-1] {
		ans = min(ans, n-p+1, p+1)
	}
	if ans == inf {
		return -1
	}
	return ans
}
