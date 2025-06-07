package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	return solve(a, b)
}

func solve(a []int, b []int) int {
	n := len(a)
	m := len(b)
	dp := make([]int, m)
	for i, j := 0, 0; i < m; i++ {
		for j < n && a[j] < b[i] {
			j++
		}
		dp[i] = j
		if j < n {
			j++
		}
	}
	if dp[m-1] < n {
		return 0
	}
	fp := make([]int, m)
	for i, j := m-1, n-1; i >= 0; i-- {
		for j >= 0 && a[j] < b[i] {
			j--
		}
		fp[i] = j
		if j >= 0 {
			j--
		}
	}

	if m == 1 {
		// 不管在哪里，只要放置b[0]即可
		return b[0]
	}

	ans := inf
	if fp[1] >= 0 {
		ans = min(ans, b[0])
	}
	if dp[m-2] < n {
		ans = min(ans, b[m-1])
	}

	for j := 0; j+2 < m; j++ {
		if dp[j] < fp[j+2] {
			ans = min(ans, b[j+1])
		}
	}

	if ans == inf {
		return -1
	}
	return ans
}

const inf = 1 << 60
