package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for range tc {
		buf.WriteString(fmt.Sprintf("%d\n", drive(reader)))
	}
	buf.WriteTo(os.Stdout)
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

func drive(reader *bufio.Reader) int {
	n, _ := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

var dd = []int{-1, 0, 1, 0, -1}
var pointer = []byte{'D', 'L', 'U', 'R'}

func solve(a []string) int {
	n := len(a)
	m := len(a[0])
	if n == 1 && m == 1 {
		return 0
	}
	deg := make([][]int, n)
	for i := range n {
		deg[i] = make([]int, m)
		for j := range m {
			if a[i][j] == '?' {
				deg[i][j] = 4
				if i == 0 {
					deg[i][j]--
				}
				if j == 0 {
					deg[i][j]--
				}
				if i == n-1 {
					deg[i][j]--
				}
				if j == m-1 {
					deg[i][j]--
				}
			} else {
				deg[i][j] = 1
			}
		}
	}

	que := make([]int, n*m)
	var head, tail int
	for i := range n {
		if a[i][0] == 'L' {
			que[head] = i * m
			head++
		}

		if a[i][m-1] == 'R' {
			que[head] = i*m + m - 1
			head++
		}
	}

	for j := range m {
		if a[0][j] == 'U' {
			que[head] = j
			head++
		}
		if a[n-1][j] == 'D' {
			que[head] = (n-1)*m + j
			head++
		}
	}

	for tail < head {
		r, c := que[tail]/m, que[tail]%m
		tail++
		for i := range 4 {
			x, y := r+dd[i], c+dd[i+1]
			if x >= 0 && x < n && y >= 0 && y < m && (a[x][y] == '?' || a[x][y] == pointer[i]) {
				deg[x][y]--
				if deg[x][y] == 0 {
					que[head] = x*m + y
					head++
				}
			}
		}
	}

	// 在队列中的，都是能离开的
	return n*m - head
}
