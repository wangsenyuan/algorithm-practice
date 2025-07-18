package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

func process(reader *bufio.Reader) bool {
	n, m := readTwoNums(reader)
	a := make([][]int, n)
	for i := range n {
		a[i] = readNNums(reader, m)
	}
	return solve(a)
}

func solve(a [][]int) bool {
	m := len(a[0])
	n := len(a)

	buf := make([][]int, n)
	for i := range n {
		buf[i] = slices.Clone(a[i])
	}

	swap := func(c1 int, c2 int) {
		for i := range n {
			buf[i][c1], buf[i][c2] = buf[i][c2], buf[i][c1]
		}
	}

	vis := make([]bool, m)
	checkRow := func(row []int) int {
		clear(vis)
		var cnt int
		for i := range m {
			if !vis[i] {
				j := i
				var c int
				for !vis[j] {
					c++
					vis[j] = true
					j = row[j] - 1
				}
				cnt += c - 1
			}
		}
		return cnt
	}

	check := func() int {
		var res int
		for i := range n {
			res = max(res, checkRow(buf[i]))
		}
		return res
	}

	tmp := check()
	if tmp <= 1 {
		return true
	}
	if tmp > 2 {
		return false
	}
	// tmp == 2

	for c1 := range m {
		for c2 := c1 + 1; c2 < m; c2++ {
			swap(c1, c2)
			if check() <= 1 {
				return true
			}
			swap(c1, c2)
		}
	}

	return false
}
