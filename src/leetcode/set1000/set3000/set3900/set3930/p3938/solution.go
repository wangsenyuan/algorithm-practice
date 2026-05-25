package p3938

func maxScore(grid [][]int) int {
	ans := -inf
	n := len(grid)

	for i, row := range grid {
		ans = max(ans, play(row, i == 0 || i == n-1))
	}
	m := len(grid[0])
	arr := make([]int, n)
	for j := range m {
		for i := range n {
			arr[i] = grid[i][j]
		}
		ans = max(ans, play(arr, j == 0 || j == m-1))
	}

	return ans
}

const inf = 1 << 60

// flag = true, 至少需要两个数字
func play(arr []int, flag bool) int {
	var sum int
	best := -inf
	var start int
	for i, x := range arr {
		sum += x
		if i-start+1 > 1 || !flag && i > 0 && i < len(arr)-1 {
			best = max(best, sum)
		}

		if start > 0 {
			best = max(best, sum+arr[start-1])
		}

		if sum < 0 {
			sum = 0
			start = i + 1
		}
	}
	return best
}
