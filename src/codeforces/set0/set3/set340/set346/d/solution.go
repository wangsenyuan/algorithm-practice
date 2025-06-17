package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, _, _ := process(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) (res string, s1 string, s2 string, virus string) {
	s1 = readString(reader)
	s2 = readString(reader)
	virus = readString(reader)
	res = solve(s1, s2, virus)
	return
}

const inf = 1 << 60

func solve(s1 string, s2 string, virus string) string {
	n := len(s1)
	m := len(s2)
	k := len(virus)
	type data struct {
		val int
		pi  int
		pj  int
		px  int
	}

	dp := make([][][]data, n+1)
	for i := range n + 1 {
		dp[i] = make([][]data, m+1)
		for j := range m + 1 {
			dp[i][j] = make([]data, k)
			for h := range k {
				dp[i][j][h] = data{-inf, 0, 0, -1}
			}
		}
	}
	dp[0][0][0].val = 0

	update := func(i int, j int, x int, i1 int, j1 int, x1 int, v int) {
		if dp[i][j][x].val < v {
			dp[i][j][x] = data{v, i1, j1, x1}
		}
	}

	pi := kmp(virus)

	for i := range n {
		for j := range m {
			for x := range k {
				if dp[i][j][x].val < 0 {
					continue
				}
				update(i+1, j, x, i, j, x, dp[i][j][x].val)
				update(i, j+1, x, i, j, x, dp[i][j][x].val)
				if s1[i] == s2[j] {
					nx := x
					for nx > 0 && s1[i] != virus[nx] {
						nx = pi[nx-1]
					}
					if s1[i] == virus[nx] {
						nx++
					}
					if nx < k {
						update(i+1, j+1, nx, i, j, x, dp[i][j][x].val+1)
					}
				}
			}
		}
	}
	i1, j1, x1 := -1, -1, -1

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for x := range k {
				if dp[i][j][x].val > 0 && (i1 < 0 || dp[i][j][x].val > dp[i1][j1][x1].val) {
					i1, j1, x1 = i, j, x
				}
			}
		}
	}
	if x1 < 0 {
		return "0"
	}
	var buf []byte

	for i1 > 0 && j1 > 0 {
		i2, j2, x2 := dp[i1][j1][x1].pi, dp[i1][j1][x1].pj, dp[i1][j1][x1].px
		if s1[i1-1] == s2[j1-1] &&
			dp[i1][j1][x1].val > 0 &&
			i1 == i2+1 && j1 == j2+1 {
			buf = append(buf, s1[i1-1])
		}
		i1, j1, x1 = i2, j2, x2
	}

	buf = reverse(buf)
	return string(buf)
}

func reverse(buf []byte) []byte {
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return buf
}

func kmp(s string) []int {
	n := len(s)
	pi := make([]int, n)
	pi[0] = 0
	for i := 1; i < n; i++ {
		j := pi[i-1]
		for j > 0 && s[i] != s[j] {
			j = pi[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}
