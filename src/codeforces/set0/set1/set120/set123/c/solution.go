package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
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
func drive(reader *bufio.Reader) string {
	n, m, k := readThreeNums(reader)
	c := make([][]int, n)
	for i := range n {
		c[i] = readNNums(reader, m)
	}
	a := solve(k, c)
	var buf bytes.Buffer
	for _, v := range a {
		buf.WriteString(fmt.Sprintf("%s\n", string(v)))
	}
	return strings.TrimSpace(buf.String())
}

type pair struct {
	first  int
	second int
}

func solve(k int, c [][]int) [][]byte {
	n := len(c)
	m := len(c[0])
	a := make([]byte, n+m-1)
	d := make([]pair, n+m-1)

	// un processed yet
	for i := range a {
		a[i] = '.'
		d[i] = pair{n * m, i}
	}
	a[0] = '('
	a[n+m-2] = ')'
	for i := range n {
		for j := range m {
			d[i+j].first = min(d[i+j].first, c[i][j])
		}
	}

	slices.SortFunc(d, func(u pair, v pair) int {
		return u.first - v.first
	})

	h := (n+m)/2 + 1

	dp := make([]int, h)
	ndp := make([]int, h)
	calc := func() int {
		clear(dp)
		dp[0] = 1

		for i := 0; i < n+m-1; i++ {
			clear(ndp)
			switch a[i] {
			case '(':
				for j := h - 1; j > 0; j-- {
					ndp[j] = min(k, dp[j-1])
				}
			case ')':
				for j := 0; j+1 < h; j++ {
					ndp[j] = min(k, dp[j+1])
				}
			default:
				for j := 0; j < h; j++ {
					if j > 0 {
						ndp[j] = min(k, dp[j-1])
					}
					if j+1 < h {
						ndp[j] = min(k, ndp[j]+dp[j+1])
					}
				}
			}
			copy(dp, ndp)
		}

		return dp[0]
	}

	for i := range n + m - 1 {
		if d[i].second == 0 || d[i].second == n+m-2 {
			continue
		}
		a[d[i].second] = '('
		// 如果这里放置左括号，能保证有k个有效的，那么就ok
		cnt := calc()
		if cnt < k {
			a[d[i].second] = ')'
			k -= cnt
		}
	}
	ans := make([][]byte, n)
	for i := range n {
		ans[i] = make([]byte, m)
		for j := range m {
			ans[i][j] = a[i+j]
		}
	}
	return ans
}
