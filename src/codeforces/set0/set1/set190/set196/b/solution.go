package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func process(reader *bufio.Reader) string {
	n, _ := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	if solve(a) {
		return "Yes"
	}
	return "No"
}

var dd = []int{-1, 0, 1, 0, -1}

const inf = 1 << 60

type pair struct {
	first  int
	second int
}

func solve(s []string) bool {
	not_visited := pair{-inf, -inf}

	pos := find(s, 'S')
	n := len(s)
	m := len(s[0])
	vis := make([][]pair, n)
	for i := range n {
		vis[i] = make([]pair, m)
		for j := range m {
			vis[i][j] = not_visited
		}
	}

	get_x := func(x int) int {
		x %= n
		if x < 0 {
			x += n
		}
		return x
	}

	get_y := func(y int) int {
		y %= m
		if y < 0 {
			y += m
		}
		return y
	}

	var que []pair
	que = append(que, pair{pos[0], pos[1]})

	for len(que) > 0 {
		it := que[0]
		que = que[1:]
		r, c := it.first, it.second
		vis[get_x(r)][get_y(c)] = it
		for i := range 4 {
			x, y := r+dd[i], c+dd[i+1]
			if s[get_x(x)][get_y(y)] != '#' {
				if vis[get_x(x)][get_y(y)] == not_visited {
					vis[get_x(x)][get_y(y)] = pair{x, y}
					que = append(que, pair{x, y})
				} else if (vis[get_x(x)][get_y(y)] != pair{x, y}) {
					return true
				}
			}
		}
	}

	return false
}

func find(s []string, c byte) []int {
	for i := range s {
		for j := range []byte(s[i]) {
			if s[i][j] == c {
				return []int{i, j}
			}
		}
	}
	return nil
}
