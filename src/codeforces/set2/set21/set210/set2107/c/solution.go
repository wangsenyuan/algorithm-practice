package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer

	tc := readNum(reader)

	for range tc {
		_, res := process(reader)
		if len(res) == 0 {
			buf.WriteString("No\n")
		} else {
			buf.WriteString("Yes\n")
			for _, x := range res {
				buf.WriteString(strconv.Itoa(x))
				buf.WriteByte(' ')
			}
			buf.WriteByte('\n')
		}
	}
	os.Stdout.Write(buf.Bytes())
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

func process(reader *bufio.Reader) (k int, res []int) {
	n, k := readTwoNums(reader)
	s := readString(reader)
	a := readNNums(reader, n)
	return k, solve(k, a, s)
}

const inf = 1e18

func solve(k int, a []int, s string) []int {
	n := len(a)
	// dp[i]表示到i为止的最大值
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			dp[i+1] = max(dp[i], 0) + a[i]
		} else {
			dp[i+1] = 0
		}
		if dp[i+1] > k {
			return nil
		}
	}

	play := func(i int, x int) []int {
		ans := make([]int, n)
		copy(ans, a)
		for j := range n {
			if s[j] == '0' {
				if j == i {
					ans[j] = x
				} else {
					ans[j] = -inf
				}
			}
		}
		return ans
	}

	var fp int
	for i := n - 1; i >= 0; i-- {
		if dp[i+1]+max(fp, 0) == k {
			// 本来就存在一个最大值为k的序列
			return play(-1, 0)
		}

		if s[i] == '0' {
			// 可以在这里设置一个数
			return play(i, k-(max(dp[i], 0)+max(fp, 0)))
		}
		fp = max(fp, 0) + a[i]
	}
	return nil
}
