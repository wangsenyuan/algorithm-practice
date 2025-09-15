package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if len(res) == 0 {
		fmt.Println("Not unique")
		return
	}
	ans := strings.Join(res, "\n")
	fmt.Println(ans)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) []string {
	n, _ := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(a []string) []string {
	n := len(a)
	m := len(a[0])
	buf := make([][]byte, n)
	for i := range n {
		buf[i] = []byte(a[i])
	}
	deg := make([]int32, n*m)
	adj := make([][]int32, n*m)

	add := func(u int, v int) {
		adj[u] = append(adj[u], int32(v))
		deg[v]++
	}

	for i := range n {
		for j := range m {
			if buf[i][j] == '.' {
				for k := range 4 {
					r, c := i+dd[k], j+dd[k+1]
					if r >= 0 && r < n && c >= 0 && c < m && buf[r][c] == '.' {
						add(i*m+j, r*m+c)
					}
				}
			}
		}
	}

	que := make([]int32, n*m)
	var head, tail int
	for i := range n {
		for j := range m {
			if buf[i][j] == '.' && deg[i*m+j] == 1 {
				que[head] = int32(i*m + j)
				head++
			}
		}
	}

	addNext := func(r int, c int) {
		for i := range 4 {
			x, y := r+dd[i], c+dd[i+1]
			if x >= 0 && x < n && y >= 0 && y < m && buf[x][y] == '.' {
				deg[x*m+y]--
				if deg[x*m+y] == 1 {
					que[head] = int32(x*m + y)
					head++
				}
			}
		}
	}

	label1 := func(r int, x int, y int) {
		if x > y {
			x, y = y, x
		}
		buf[r][x] = '<'
		buf[r][y] = '>'
	}

	label2 := func(c int, x int, y int) {
		if x > y {
			x, y = y, x
		}
		buf[x][c] = '^'
		buf[y][c] = 'v'
	}

	for tail < head {
		r, c := int(que[tail])/m, int(que[tail])%m
		tail++
		// deg[r][c] = 1
		for _, nxt := range adj[r*m+c] {
			x, y := int(nxt)/m, int(nxt)%m
			if buf[x][y] == '.' {
				if x == r {
					// 水平
					label1(r, c, y)
				} else {
					label2(c, r, x)
				}
				addNext(x, y)
				break
			}
		}
	}

	for i := range n {
		a[i] = string(buf[i])
		for j := range m {
			if buf[i][j] == '.' {
				return nil
			}
		}
	}

	return a
}
