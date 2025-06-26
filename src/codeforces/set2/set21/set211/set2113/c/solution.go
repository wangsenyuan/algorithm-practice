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
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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
	g := make([]string, n)
	for i := range n {
		g[i] = readString(reader)
	}
	return solve(k, g)
}

func solve(k int, g []string) int {
	n := len(g)
	m := len(g[0])
	sum := make([][]int, n)
	for i := range n {
		sum[i] = make([]int, m)
		for j := range m {
			if g[i][j] == 'g' {
				sum[i][j]++
			}
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}
	all := sum[n-1][m-1]
	var best int

	for i := range n {
		for j := range m {
			if g[i][j] == '.' {
				r0 := max(0, i-k+1)
				c0 := max(0, j-k+1)
				r1 := min(n-1, i+k-1)
				c1 := min(m-1, j+k-1)
				tmp := sum[r1][c1]
				if r0 > 0 {
					tmp -= sum[r0-1][c1]
				}
				if c0 > 0 {
					tmp -= sum[r1][c0-1]
				}
				if r0 > 0 && c0 > 0 {
					tmp += sum[r0-1][c0-1]
				}
				best = max(best, all-tmp)
			}
		}
	}

	return best
}
