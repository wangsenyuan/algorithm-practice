package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []string) bool {
	n := len(a)
	m := len(a[0])

	var corners [][]int
	for i := range n {
		for j := range m {
			if a[i][j] == '1' {
				corners = append(corners, []int{i, j})
			}
		}
	}
	if len(corners) != 4 {
		return false
	}

	if corners[0][0] != corners[1][0] || corners[2][0] != corners[3][0] {
		return false
	}

	if corners[0][1] != corners[2][1] || corners[1][1] != corners[3][1] {
		return false
	}

	w := corners[1][1] - corners[0][1]
	h := corners[2][0] - corners[0][0]
	if w < 2 || h < 2 {
		return false
	}

	t := corners[0][0]
	r := corners[1][1]
	b := corners[2][0]
	l := corners[0][1]

	if w == 2 || h == 2 {
		for i := range n {
			for j := range m {
				if i < t || j < l || i > b || j > r {
					if a[i][j] != '0' {
						return false
					}
				} else if i == t || i == b || j == l || j == r {
					if a[i][j] != '1' && a[i][j] != '2' {
						return false
					}
				} else if a[i][j] != '4' {
					return false
				}
			}
		}
		return true
	}

	for i := range n {
		for j := range m {
			var g int
			if i < t || i > b || j < l || j > r {
				g = 0
			} else {
				if i == t || i == b {
					if j == r || j == l {
						g = 1
					} else {
						g = 2
					}
				} else if j == l || j == r {
					g = 2
				} else if i > t && i < b && j > l && j < r {
					g = 4
				}
			}

			if g != int(a[i][j]-'0') {
				return false
			}
		}
	}

	return true
}
