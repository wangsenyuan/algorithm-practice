package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(len(res))
	if len(res) > 0 {
		s := fmt.Sprintf("%v", res)
		fmt.Println(s[1 : len(s)-1])
	}
}

func process(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	image := make([]string, n)
	for i := range n {
		var buf []byte
		for range m {
			var x int
			fmt.Fscan(reader, &x)
			buf = append(buf, byte(x+'0'))
		}
		image[i] = string(buf)
	}
	return solve(image)
}

func solve(image []string) []int {
	n := len(image)
	m := len(image[0])
	a := make([][]byte, n)
	b := make([][]byte, n)
	c := make([][]byte, n)
	for i := 0; i < n; i++ {
		a[i] = []byte(image[i])
		b[i] = []byte(image[i])
		c[i] = make([]byte, m)
	}

	check := func(i int, j int) bool {
		if i >= 0 && i < n && j >= 0 && j < m {
			return b[i][j] == '1'
		}
		return false
	}

	var dd = []int{-1, 0, 1, 0, -1}

	erosion := func() {
		for i := range n {
			copy(c[i], b[i])
		}

		for i := 1; i < n-1; i++ {
			for j := 1; j < m-1; j++ {
				for u := 0; u < 4; u++ {
					x, y := i+dd[u], j+dd[u+1]
					if !check(x, y) {
						c[i][j] = '0'
						break
					}
				}
			}
		}

		for i := range n {
			copy(b[i], c[i])
		}
	}

	for range 4 {
		erosion()
	}

	delation := func() {
		for i := range n {
			copy(c[i], b[i])
		}

		for i := 1; i < n-1; i++ {
			for j := 1; j < m-1; j++ {
				for u := 0; u < 4; u++ {
					x, y := i+dd[u], j+dd[u+1]
					if check(x, y) {
						c[i][j] = '1'
						break
					}
				}
			}
		}

		for i := range n {
			copy(b[i], c[i])
		}
	}

	for range 8 {
		delation()
	}

	for i := range n {
		for j := range m {
			if b[i][j] == '1' {
				a[i][j] = '0'
			}
		}
	}

	que2 := make([]int, n*m)
	bfs2 := func(x int, y int) int {
		var head, tail int
		que2[head] = x*m + y
		a[x][y] = '0'
		head++
		for tail < head {
			r, c := que2[tail]/m, que2[tail]%m
			tail++
			for u := range 4 {
				x, y := r+dd[u], c+dd[u+1]
				if x >= 0 && x < n && y >= 0 && y < m && a[x][y] == '1' {
					a[x][y] = '0'
					que2[head] = x*m + y
					head++
				}
			}
		}
		if head > 5 {
			return 1
		}
		return 0
	}

	que := make([]int, n*m)

	bfs := func(x int, y int) int {
		b[x][y] = '0'
		var head, tail int
		que[head] = x*m + y
		head++
		var sz int
		for tail < head {
			r, c := que[tail]/m, que[tail]%m
			tail++
			for u := range 4 {
				x, y := r+dd[u], c+dd[u+1]
				if x >= 0 && x < n && y >= 0 && y < m {
					if a[x][y] == '1' {
						sz += bfs2(x, y)
					}
					if b[x][y] == '1' {
						b[x][y] = '0'
						que[head] = x*m + y
						head++
					}
				}
			}
		}

		return sz
	}

	var ans []int

	for i := range n {
		for j := range m {
			if b[i][j] == '1' {
				ans = append(ans, bfs(i, j))
			}
		}
	}

	sort.Ints(ans)
	return ans
}
