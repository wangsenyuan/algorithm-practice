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
	n, m, k := readThreeNums(reader)
	forbidden := make([]string, k)
	for i := range k {
		forbidden[i] = readString(reader)
	}
	return solve(n, m, forbidden)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

type mat [][]int

func (a mat) mul(b mat) mat {
	n := len(a)
	m := len(b)
	k := len(b[0])
	c := make(mat, n)
	for i := range c {
		c[i] = make([]int, k)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			for u := 0; u < m; u++ {
				c[i][j] = add(c[i][j], mul(a[i][u], b[u][j]))
			}
		}
	}
	return c
}

func identity(n int) mat {
	res := make([][]int, n)
	for i := range n {
		res[i] = make([]int, n)
		res[i][i] = 1
	}
	return res
}

func pow(a mat, n int) mat {
	if n == 0 {
		return identity(len(a))
	}
	res := pow(a, n>>1)
	res = res.mul(res)
	if n&1 == 1 {
		res = a.mul(res)
	}
	return res
}

func code(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a')
	}
	return 26 + int(c-'A')
}

func solve(n int, m int, forbidden []string) int {
	a := make(mat, m)
	for i := range a {
		a[i] = make([]int, m)
		for j := range m {
			a[i][j] = 1
		}
	}
	for _, cur := range forbidden {
		x, y := code(cur[0]), code(cur[1])
		a[x][y] = 0
	}
	a = pow(a, n-1)
	v := make(mat, m)
	for i := range m {
		v[i] = make([]int, 1)
		v[i][0] = 1
	}
	res := a.mul(v)
	var ans int
	for i := range m {
		ans = add(ans, res[i][0])
	}
	return ans
}
