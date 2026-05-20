package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t, k int
	fmt.Fscan(reader, &t, &k)
	queries := make([]string, t)
	for i := range t {
		var l, r string
		fmt.Fscan(reader, &l, &r)
		queries[i] = l + " " + r
	}

	res := solve(k, queries)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, ans := range res {
		fmt.Fprintln(writer, ans)
	}
}

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func solve(k int, queries []string) []int {
	maxLen := 0
	arr := make([][2]string, len(queries))
	for i, cur := range queries {
		j := strings.Index(cur, " ")
		l, r := cur[:j], cur[j+1:]
		arr[i] = [2]string{l, r}
		maxLen = max(maxLen, len(l), len(r))
	}

	counter := newCounter(k, maxLen)
	ans := make([]int, len(queries))
	for i, cur := range arr {
		l, r := cur[0], cur[1]
		ans[i] = sub(counter.countNearly(r), counter.countNearly(prev(l)))
	}
	return ans
}

type counter struct {
	k    int
	free [][]int
}

func newCounter(k int, maxLen int) counter {
	free := make([][]int, maxLen+1)
	for i := range free {
		free[i] = make([]int, k+2)
	}
	for d := 1; d <= k+1; d++ {
		free[0][d] = 1
	}
	for rem := 1; rem <= maxLen; rem++ {
		for d := k + 1; d >= 1; d-- {
			nd := min(k+1, d+1)
			free[rem][d] = 8 * free[rem-1][nd] % mod
			if d > k {
				free[rem][d] = add(free[rem][d], 2*free[rem-1][1]%mod)
			}
		}
	}
	return counter{k, free}
}

func (c counter) countNearly(s string) int {
	if s == "0" {
		return 0
	}
	total := value(s)
	bad := c.countSparseLucky(s)
	return sub(total, bad)
}

func (c counter) countSparseLucky(s string) int {
	n := len(s)
	res := 0
	for length := 1; length < n; length++ {
		res = add(res, 7*c.free[length-1][c.k+1]%mod)
		res = add(res, 2*c.free[length-1][1]%mod)
	}

	dist := c.k + 1
	for i := 0; i < n; i++ {
		limit := int(s[i] - '0')
		start := 0
		if i == 0 {
			start = 1
		}
		rem := n - i - 1
		for x := start; x < limit; x++ {
			if isLucky(x) {
				if dist > c.k {
					res = add(res, c.free[rem][1])
				}
			} else {
				res = add(res, c.free[rem][min(c.k+1, dist+1)])
			}
		}
		if isLucky(limit) {
			if dist <= c.k {
				return res
			}
			dist = 1
		} else {
			dist = min(c.k+1, dist+1)
		}
	}
	return add(res, 1)
}

func isLucky(x int) bool {
	return x == 4 || x == 7
}

func value(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		res = (res*10 + int(s[i]-'0')) % mod
	}
	return res
}

func prev(s string) string {
	buf := []byte(s)
	i := len(buf) - 1
	for i >= 0 && buf[i] == '0' {
		buf[i] = '9'
		i--
	}
	if i >= 0 {
		buf[i]--
	}
	for len(buf) > 1 && buf[0] == '0' {
		buf = buf[1:]
	}
	return string(buf)
}
