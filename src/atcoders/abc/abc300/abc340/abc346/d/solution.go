package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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
	n := readNum(reader)
	s := readString(reader)
	a := readNNums(reader, n)
	return solve(s, a)
}

const inf = 1 << 60

func solve(s string, a []int) int {
	n := len(s)
	dp := make([][2]int, n+1)
	for i := 0; i < n; i++ {
		x := int(s[i] - '0')
		// 修改当前位，或者不修改
		dp[i+1][x] = dp[i][x^1]
		dp[i+1][x^1] = dp[i][x] + a[i]
	}
	best := inf

	fp := make([]int, 2)

	for i := n - 1; i > 0; i-- {
		x := int(s[i] - '0')
		// 当前位为x，前面一位也为x时
		best = min(best, dp[i][x]+fp[x^1])
		best = min(best, dp[i][x^1]+a[i]+fp[x])

		fp[x^1], fp[x] = fp[x]+a[i], fp[x^1]
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
