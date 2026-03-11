package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) bool {
	first := readString(reader)
	ss := strings.Split(first, " ")
	n, _ := strconv.Atoi(ss[0])

	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

func solve(a []string) bool {
	n := len(a)
	m := len(a[0])

	var cells [][]int
	for i := range n {
		for j := range m {
			if a[i][j] == 'B' {
				cells = append(cells, []int{i, j})
			}
		}
	}

	dp := make([][][]int, n)
	for i := range n {
		dp[i] = make([][]int, m)
		for j := range m {
			dp[i][j] = make([]int, 4)
		}
	}

	for i := range n {
		for j := range m {
			if a[i][j] == 'B' {
				if i > 0 {
					dp[i][j][0] = dp[i-1][j][0] + 1
				}
				if j > 0 {
					dp[i][j][1] = dp[i][j-1][1] + 1
				}
			} else {
				dp[i][j][0] = -1
				dp[i][j][1] = -1
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if a[i][j] == 'B' {
				if i < n-1 {
					dp[i][j][2] = dp[i+1][j][2] + 1
				}
				if j < m-1 {
					dp[i][j][3] = dp[i][j+1][3] + 1
				}
			} else {
				dp[i][j][2] = -1
				dp[i][j][3] = -1
			}
		}
	}

	check := func(a []int, b []int) bool {
		if a[0] > b[0] {
			a, b = b, a
		}
		// 0 for top, 1 for left, 2 for bottom, 3 for right
		// a[0] <= b[0]
		if a[1] <= b[1] {
			if dp[a[0]][b[1]][1] >= b[1]-a[1] && dp[a[0]][b[1]][2] >= b[0]-a[0] {
				return true
			}
			if dp[b[0]][a[1]][0] >= b[0]-a[0] && dp[b[0]][a[1]][3] >= b[1]-a[1] {
				return true
			}
		} else {
			// b在a左边
			if dp[a[0]][b[1]][3] >= a[1]-b[1] && dp[a[0]][b[1]][2] >= b[0]-a[0] {
				return true
			}
			if dp[b[0]][a[1]][0] >= b[0]-a[0] && dp[b[0]][a[1]][1] >= a[1]-b[1] {
				return true
			}
		}

		return false
	}

	for i := range cells {
		for j := range cells {
			if !check(cells[i], cells[j]) {
				return false
			}
		}
	}
	return true
}
