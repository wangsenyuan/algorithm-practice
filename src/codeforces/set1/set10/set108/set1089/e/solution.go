package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	res := solve(n)
	fmt.Println(strings.Join(res, " "))
}

func solve(n int) []string {
	path := buildPath()
	m := len(path)

	reach := make([][]bool, m)
	pre := make([][]int, m)
	for i := 0; i < m; i++ {
		reach[i] = make([]bool, m+1)
		pre[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			pre[i][j] = -1
		}
	}

	reach[0][1] = true
	for i := 1; i < m; i++ {
		for j := 0; j < i; j++ {
			if !sameLine(path[i], path[j]) {
				continue
			}
			for l := 2; l <= n+1; l++ {
				if reach[j][l-1] && !reach[i][l] {
					reach[i][l] = true
					pre[i][l] = j
				}
			}
		}
	}

	cur := m - 1
	l := n + 1
	ans := make([]string, 0, l)
	for cur >= 0 && l > 0 {
		ans = append(ans, cellName(path[cur]))
		cur = pre[cur][l]
		l--
	}
	reverse(ans)
	return ans
}

func buildPath() []int {
	var path []int
	for col := 0; col < 6; col++ {
		if col%2 == 0 {
			for row := 0; row < 8; row++ {
				path = append(path, col*8+row)
			}
		} else {
			for row := 7; row >= 0; row-- {
				path = append(path, col*8+row)
			}
		}
	}

	// A Hamiltonian path on the last two columns, from g1 to h8.
	tail := []string{
		"g1", "g2", "g3", "g4", "g5", "g6", "g8", "g7",
		"h7", "h1", "h2", "h3", "h4", "h5", "h6", "h8",
	}
	for _, s := range tail {
		path = append(path, parseCell(s))
	}

	return path
}

func sameLine(a, b int) bool {
	return a/8 == b/8 || a%8 == b%8
}

func cellName(x int) string {
	return fmt.Sprintf("%c%d", 'a'+x/8, x%8+1)
}

func parseCell(s string) int {
	col := int(s[0] - 'a')
	row := int(s[1] - '1')
	return col*8 + row
}

func reverse[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
