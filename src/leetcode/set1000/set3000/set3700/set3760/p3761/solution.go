package p3761

func minMirrorPairDistance(nums []int) int {
	n := len(nums)
	pos := make(map[int]int)

	ans := n

	for i := range n {
		if j, ok := pos[nums[i]]; ok {
			ans = min(ans, i-j)
		}
		w := reverse(nums[i])
		pos[w] = i
	}
	if ans == n {
		return -1
	}
	return ans
}

func reverse(num int) int {
	var res int
	for num > 0 {
		x := num % 10
		res = res*10 + x
		num /= 10
	}
	return res
}
