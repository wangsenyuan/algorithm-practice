package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	n, _ := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

func solve(a []string) int {
	n := len(a)
	m := len(a[0])

	dp := make([][]int, n)

	for i, s := range a {
		dp[i] = make([]int, m)
		prev := -inf
		for j := range 2 * m {
			if s[j%m] == '1' {
				prev = j
			}
			dp[i][j%m] = j - prev
		}
		// 这一行没有1
		if prev < 0 {
			return -1
		}
		next := inf
		for j := 2*m - 1; j >= 0; j-- {
			if s[j%m] == '1' {
				next = j
			}
			dp[i][j%m] = min(dp[i][j%m], next-j)
		}
	}
	best := inf
	for j := range m {
		var sum int
		for i := range n {
			sum += dp[i][j]
		}
		best = min(best, sum)
	}

	return best
}

const inf = 1 << 60
