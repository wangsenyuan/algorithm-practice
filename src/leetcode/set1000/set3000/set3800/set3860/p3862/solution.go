package p3862

func smallestBalancedIndex(nums []int) int {
	n := len(nums)
	pref := make([]int, n+1)
	for i := range n {
		if i == 0 {
			pref[i] = nums[i]
		} else {
			pref[i] = nums[i] + pref[i-1]
		}
	}
	ans := -1
	cur := 1
	for i := n - 1; i > 0; i-- {
		if pref[i-1] == cur {
			ans = i
		}
		if pref[i-1]/nums[i] < cur || pref[i-1] < cur*nums[i] {
			break
		}
		cur *= nums[i]
	}
	return ans
}
