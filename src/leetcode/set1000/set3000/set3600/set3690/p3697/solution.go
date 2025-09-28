package p3697

func splitArray(nums []int) int64 {
	n := len(nums)

	ok := make([]bool, n+1)
	suf := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		if i+1 < n && nums[i] <= nums[i+1] {
			break
		}
		ok[i] = true
		suf[i] = suf[i+1] + nums[i]
	}

	ans := -1
	var pref int

	for i := 0; i+1 < n; i++ {
		if i > 0 && nums[i] <= nums[i-1] {
			break
		}
		pref += nums[i]
		tmp := abs(pref - suf[i+1])
		if ok[i+1] && (ans < 0 || tmp < ans) {
			ans = tmp
		}
	}

	return int64(ans)
}

func abs(num int) int {
	return max(num, -num)
}
