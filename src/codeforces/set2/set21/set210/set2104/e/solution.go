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
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) []int {
	_, k := readTwoNums(reader)
	s := readString(reader)
	m := readNum(reader)
	queries := make([]string, m)
	for i := range m {
		queries[i] = readString(reader)
	}
	return solve(k, s, queries)
}

func solve(k int, s string, queries []string) []int {
	n := len(s)
	dp := make([]int, n+2)
	dp[n+1] = 0
	dp[n] = 1
	next := make([][]int, n+1)
	for i := range n + 1 {
		next[i] = make([]int, k)
	}
	for i := range k {
		next[n][i] = n
	}
	for i := n - 1; i >= 0; i-- {
		x := int(s[i] - 'a')
		// 如果在i处开始匹配，那么把后缀匹配完后，再加一个字符即可
		copy(next[i], next[i+1])
		next[i][x] = i
		dp[i] = dp[i+1] + 1
		for y := 0; y < k; y++ {
			dp[i] = min(dp[i], dp[next[i][y]+1]+1)
		}
	}

	find := func(t string) int {
		var j int
		for i := 0; i < len(t); i++ {
			x := int(t[i] - 'a')
			j = next[j][x]
			if j == n {
				return 0
			}
			j++
		}
		return dp[j]
	}

	res := make([]int, len(queries))

	for i, t := range queries {
		res[i] = find(t)
	}

	return res
}
