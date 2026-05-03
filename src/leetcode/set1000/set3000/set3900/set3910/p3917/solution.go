package p3917

func countOppositeParity(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	cnt := make([]int, 2)
	for i := n - 1; i >= 0; i-- {
		ans[i] = cnt[(nums[i]&1)^1]
		cnt[nums[i]&1]++
	}
	return ans
}
