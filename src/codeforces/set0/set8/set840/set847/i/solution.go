package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := drive(reader)
	fmt.Println(ans)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	first := readNNums(reader, 4)
	n := first[0]
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(first[2], first[3], a)
}

func solve(q int, p int, a []string) int {
	n := len(a)
	m := len(a[0])

	val := make([][]int, n)
	marked := make([][]int, n)
	for i := range n {
		val[i] = make([]int, m)
		marked[i] = make([]int, m)
		for j := range m {
			marked[i][j] = -1
		}
	}

	que := make([]int, n*m)

	var dd = []int{-1, 0, 1, 0, -1}

	bfs := func(r int, c int, w int) {
		var head, tail int
		marked[r][c] = w
		que[head] = r*m + c
		head++
		for tail < head {
			x, y := que[tail]/m, que[tail]%m
			tail++
			if marked[x][y] <= 1 {
				continue
			}
			for i := range 4 {
				u, v := x+dd[i], y+dd[i+1]
				if u >= 0 && u < n && v >= 0 && v < m && marked[u][v] == -1 && a[u][v] != '*' {
					marked[u][v] = marked[x][y] / 2
					que[head] = u*m + v
					head++
				}
			}
		}
		for i := range head {
			x, y := que[i]/m, que[i]%m
			val[x][y] += marked[x][y]
			marked[x][y] = -1
		}
	}

	for i := range n {
		for j := range m {
			if a[i][j] >= 'A' && a[i][j] <= 'Z' {
				x := int(a[i][j] - 'A')
				bfs(i, j, (x+1)*q)
			}
		}
	}
	var ans int
	for i := range n {
		for j := range m {
			if val[i][j] > p {
				ans++
			}
		}
	}

	return ans
}
