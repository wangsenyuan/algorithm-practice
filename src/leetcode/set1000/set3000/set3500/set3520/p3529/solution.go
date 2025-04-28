package p3529

func countCells(grid [][]byte, pattern string) int {
	n := len(grid)
	m := len(grid[0])

	p := kmp(pattern)
	k := len(pattern)

	row := make([]int, n*m+1)

	for i, j := 0, 0; i < n*m; i++ {
		for j > 0 && pattern[j] != grid[i/m][i%m] {
			j = p[j-1]
		}
		if pattern[j] == grid[i/m][i%m] {
			j++
		}
		if j == k {
			row[i-k+1]++
			row[i+1]--
			j = p[j-1]
		}
	}
	for i := range n * m {
		if i > 0 {
			row[i] += row[i-1]
		}
	}
	col := make([]int, n*m+1)
	for i, j := 0, 0; i < n*m; i++ {
		for j > 0 && pattern[j] != grid[i%n][i/n] {
			j = p[j-1]
		}
		if pattern[j] == grid[i%n][i/n] {
			j++
		}
		if j == k {
			col[i-k+1]++
			col[i+1]--
			j = p[j-1]
		}
	}

	for i := range n * m {
		if i > 0 {
			col[i] += col[i-1]
		}
	}
	var ans int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if row[i*m+j] > 0 && col[j*n+i] > 0 {
				ans++
			}
		}
	}
	return ans
}

func kmp(s string) []int {
	n := len(s)
	p := make([]int, n)
	for i := 1; i < n; i++ {
		j := p[i-1]
		for j > 0 && s[i] != s[j] {
			j = p[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		p[i] = j
	}
	return p
}
