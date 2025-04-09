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
	ans := process(reader)
	var buf bytes.Buffer
	for _, x := range ans {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) []int {
	s := readString(reader)
	m := readNum(reader)
	qs := make([][]int, m)
	for i := range m {
		qs[i] = readNNums(reader, 2)
	}
	return solve(s, qs)
}

type pair struct {
	first  int
	second int
}

func solve(s string, queries [][]int) []int {
	n := len(s)

	at := make([][]pair, n)
	for i, cur := range queries {
		l, r := cur[0], cur[1]
		r--
		l--
		at[r] = append(at[r], pair{l, i})
	}

	pal := make([][]bool, n)
	for i := range n {
		pal[i] = make([]bool, n)
		pal[i][i] = true
	}
	dp := make([]int, n)
	pref := make([]int, n+1)
	add := make([]int, n+1)
	cross := make([]int, n+1)
	ans := make([]int, len(queries))
	for r := range n {
		pref[r+1] = pref[r]
		for l := r; l >= 0; l-- {
			if l == r || s[l] == s[r] && (l+1 == r || pal[l+1][r-1]) {
				pal[l][r] = true
			}
			if pal[l][r] {
				pref[r+1]++
				add[l]++
				dp[l]++
			}
		}
		// 对于l...r必须知道，那些中点在和 l相交的区域
		for l := 0; l <= r; l++ {
			if l > 0 {
				add[l] += add[l-1]
			}
			cross[l] += add[l]
		}

		for _, cur := range at[r] {
			l, i := cur.first, cur.second
			ans[i] = pref[r+1] - pref[l]
			ans[i] -= cross[l]
			// dp[l]是从l开始的回文串的个数， 这部分会在cross中被减掉
			ans[i] += dp[l]
		}
		clear(add[:r+1])
	}
	return ans
}
