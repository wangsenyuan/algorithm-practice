package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, grid := range drive(reader) {
		for _, row := range grid {
			s := fmt.Sprintf("%v", row)
			fmt.Fprintln(writer, s[1:len(s)-1])
		}
	}
}

func drive(reader *bufio.Reader) [][][]int {
	var tc int
	fmt.Fscan(reader, &tc)
	res := make([][][]int, tc)
	for i := range tc {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		g := make([]string, n)
		for j := range n {
			fmt.Fscan(reader, &g[j])
		}
		res[i] = solve(g)
	}
	return res
}

const inf = 1 << 60

func solve(g []string) [][]int {
	n, m := len(g), len(g[0])
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			if g[i][j] == '1' {
				a[i][j] = 1
			}
		}
	}

	transposed := false
	if n > m {
		a = transpose(a)
		n, m = m, n
		transposed = true
	}

	ans := make([][]int, n)
	for i := range n {
		ans[i] = make([]int, m)
		for j := range m {
			ans[i][j] = inf
		}
	}

	smin := make([][]int, n)
	for i := range n {
		smin[i] = make([]int, m)
	}
	left := make([]int, m)
	right := make([]int, m)

	for top := range n {
		for i := range n {
			for j := range m {
				smin[i][j] = inf
			}
		}
		for bot := top + 1; bot < n; bot++ {
			clear(left)
			clear(right)
			for j := range m {
				if a[top][j] == 1 && a[bot][j] == 1 {
					left[j] = j + 1
				} else if j > 0 {
					left[j] = left[j-1]
				}
			}
			for j := m - 2; j >= 0; j-- {
				if a[top][j+1] == 1 && a[bot][j+1] == 1 {
					right[j] = j + 2
				} else {
					right[j] = right[j+1]
				}
			}
			h := bot - top + 1
			for j := range m {
				if a[top][j] == 1 && a[bot][j] == 1 {
					if j > 0 && left[j-1] > 0 {
						smin[bot][j] = min(smin[bot][j], (j+1-left[j-1]+1)*h)
					}
					if right[j] > 0 {
						smin[bot][j] = min(smin[bot][j], (right[j]-j)*h)
					}
				} else if left[j] > 0 && right[j] > 0 {
					smin[bot][j] = min(smin[bot][j], (right[j]-left[j]+1)*h)
				}
			}
		}
		for j := range m {
			ans[n-1][j] = min(ans[n-1][j], smin[n-1][j])
		}
		for bot := n - 2; bot >= top; bot-- {
			for j := range m {
				smin[bot][j] = min(smin[bot][j], smin[bot+1][j])
				ans[bot][j] = min(ans[bot][j], smin[bot][j])
			}
		}
	}

	for i := range n {
		for j := range m {
			if ans[i][j] == inf {
				ans[i][j] = 0
			}
		}
	}

	if transposed {
		return transpose(ans)
	}
	return ans
}

func transpose(a [][]int) [][]int {
	n, m := len(a), len(a[0])
	b := make([][]int, m)
	for j := range m {
		b[j] = make([]int, n)
		for i := range n {
			b[j][i] = a[i][j]
		}
	}
	return b
}
