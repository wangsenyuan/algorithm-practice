package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, _, k := readThreeNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(k, a)
}

func rotate(a []string) []string {
	n := len(a)
	m := len(a[0])
	buf := make([][]byte, m)
	for i := range m {
		buf[i] = make([]byte, n)
	}
	for i := range n {
		for j := range m {
			buf[j][i] = a[i][j]
		}
	}
	res := make([]string, m)
	for i := range m {
		res[i] = string(buf[i])
	}
	return res
}
func solve(k int, a []string) int {
	if len(a) > len(a[0]) {
		a = rotate(a)
	}
	n := len(a)
	m := len(a[0])

	cols := make([][]int, m)
	for i := range m {
		cols[i] = make([]int, n+1)
	}

	isCross := func(i int, j int) bool {
		if a[i][j] == '0' {
			return false
		}
		if i == 0 || a[i-1][j] == '0' || i == n-1 || a[i+1][j] == '0' {
			return false
		}
		if j == 0 || a[i][j-1] == '0' || j == m-1 || a[i][j+1] == '0' {
			return false
		}
		return true
	}

	for i := range n {
		for j := range m {
			cols[j][i+1] = cols[j][i]
			if isCross(i, j) {
				cols[j][i+1]++
			}
		}
	}

	get := func(c int, r1 int, r2 int) int {
		// r1 < r2 holds
		tmp := cols[c][r2+1] - cols[c][r1]
		if isCross(r1, c) {
			tmp--
		}
		if r1 != r2 && isCross(r2, c) {
			tmp--
		}
		return tmp
	}

	var res int

	// n * n * m
	for r1 := range n {
		for r2 := r1; r2 < n; r2++ {
			var sum int
			var prev int
			for c1, c2 := 0, 0; c2 < m; c2++ {
				sum += get(c2, r1, r2)
				for sum-get(c1, r1, r2) >= k {
					sum -= get(c1, r1, r2)
					c1++
				}
				if sum >= k {
					res += (c1 - prev) * (m - c2 - 1)
					prev = c1
					sum -= get(c1, r1, r2)
					c1++
				}
			}
		}
	}

	return res
}
