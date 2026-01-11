package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	first := readString(reader)
	ss := strings.Split(first, " ")
	n, _ := strconv.Atoi(ss[0])
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
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

func solve(a []string) int {
	if len(a) > len(a[0]) {
		a = rotate(a)
	}
	n := len(a)
	// n <= 16
	m := len(a[0])
	dp := make([][2]int, 1<<n)
	ndp := make([][2]int, 1<<n)
	dp[0][0] = 1

	var play func(prev int, r int, j int, c int, v int, mask int, top int)

	play = func(prev int, r int, j int, c int, v int, mask int, top int) {
		if r == n {
			ndp[mask][c] = add(ndp[mask][c], v)
			return
		}
		if a[r][j] == 'x' {
			play(prev, r+1, j, c, v, mask, -1)
			return
		}
		// 如果这里放置
		play(prev, r+1, j, c, v, mask|(1<<r), r)
		// 考虑这里不放置

		if (prev>>r)&1 == 0 && top == -1 && c == 0 {
			play(prev, r+1, j, c+1, v, mask, top)
		}

		if (prev>>r)&1 == 1 || top != -1 {
			// 被左边，或者是上边已经覆盖到了，可以不放置
			if (prev>>r)&1 == 1 {
				mask |= 1 << r
			}
			play(prev, r+1, j, c, v, mask, top)
		}
	}

	for i := range m {
		for mask := range 1 << n {
			for c := range 2 {
				if dp[mask][c] != 0 {
					play(mask, 0, i, c, dp[mask][c], 0, -1)
				}
			}
		}
		for s := range 1 << n {
			for c := range 2 {
				dp[s][c] = ndp[s][c]
				ndp[s][c] = 0
			}
		}
	}

	var res int
	for s := range 1 << n {
		for c := range 2 {
			res = add(res, dp[s][c])
		}
	}
	return res
}

func rotate(a []string) []string {
	n := len(a)
	m := len(a[0])
	res := make([][]byte, m)
	for i := range m {
		res[i] = make([]byte, n)
	}
	for i := range n {
		for j := range m {
			res[j][i] = a[i][j]
		}
	}
	b := make([]string, m)
	for i := range m {
		b[i] = string(res[i])
	}
	return b
}
