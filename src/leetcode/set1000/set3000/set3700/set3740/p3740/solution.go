package p3740

func minimumDistance(nums []int) int {
	// i < j < k
	// abs(j - i) + abs(k - j) + abs(k - i)
	n := len(nums)

	res := -1

	pos := make([][]int, n+1)

	for i, v := range nums {
		pos[v] = append(pos[v], i)
		for len(pos[v]) > 3 {
			pos[v] = pos[v][1:]
		}
		if len(pos[v]) == 3 {
			cur := pos[v][2] - pos[v][0]
			cur *= 2
			if res == -1 || cur < res {
				res = cur
			}
		}
	}

	return res
}
