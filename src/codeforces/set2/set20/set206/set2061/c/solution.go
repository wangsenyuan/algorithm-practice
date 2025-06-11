package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
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
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(a []int) int {
	n := len(a)
	if n == 1 {
		if a[0] == 0 {
			return 2
		}
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	// dp[i]表示第i个人说真话时的计数
	for i := 1; i <= n; i++ {
		if a[i-1] < i {
			// i如果说真话
			if i == 1 {
				dp[i] = 1
				continue
			}
			// i >= 2

			if a[i-2] == a[i-1] {
				dp[i] = dp[i-1]
			}
			if i == 2 && a[i-1] == 1 || i >= 3 && a[i-3]+1 == a[i-1] {
				// 在i说真话的情况下，且i-1是在说假话, 但是i-2说的是真话
				dp[i] = add(dp[i], dp[i-2])
			}
		}
	}
	// 要么n在说假话，要么说真话
	return add(dp[n-1], dp[n])
}

func solve1(a []int) int {
	n := len(a)
	if n == 1 {
		if a[0] == 0 {
			return 2
		}
		return 1
	}
	type pair struct {
		first  int
		second int
	}
	// dp[i]是它说真话时的状态
	dp := make([]map[int]int, n+1)
	for i := range n + 1 {
		dp[i] = make(map[int]int)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		// 如果i说的是假话
		if i+1 <= n && a[i] > 0 {
			dp[i+1][a[i]] = add(dp[i+1][a[i]], dp[i-1][a[i]-1])
		}
		// 如果i说的是真话
		if a[i-1] < i {
			dp[i][a[i-1]] = add(dp[i][a[i-1]], dp[i-1][a[i-1]])
		}
	}
	// 最后一个人说的是真话
	res := dp[n][a[n-1]]
	if n >= 2 {
		// 最后一个人说假话，但是倒数第二个人说真话
		res = add(res, dp[n-1][a[n-2]])
	}
	return res
}
