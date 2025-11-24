package p3752

import "slices"

func lexSmallestNegatedPerm(n int, target int64) []int {
	sum := int64((1 + n) * n / 2)
	if target > sum || target < -sum || abs(sum-target)%2 == 1 {
		return nil
	}
	ans := make([]int, n)
	for i := range n {
		ans[i] = i + 1
	}

	for i := n; i > 0; i-- {
		if sum-target >= 2*int64(i) {
			ans[i-1] *= -1
			sum -= 2 * int64(i)
		}
	}
	slices.Sort(ans)
	return ans
}

func abs(num int64) int64 {
	return max(num, -num)
}
