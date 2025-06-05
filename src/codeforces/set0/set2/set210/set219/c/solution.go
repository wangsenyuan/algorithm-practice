package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans, res, _ := process(reader)
	fmt.Println(ans)
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

func process(reader *bufio.Reader) (ans int, res string, s string) {
	_, k := readTwoNums(reader)
	s = readString(reader)
	ans, res = solve(k, s)
	return
}

type pair struct {
	first  int
	second int
}

func code(c byte) int {
	return int(c - 'A')
}

func solve(k int, s string) (int, string) {
	n := len(s)

	dp := make([][]pair, n)
	for i := range n {
		dp[i] = make([]pair, k)
		for j := range k {
			dp[i][j] = pair{n, -1}
		}
	}

	for i := range k {
		dp[0][i] = pair{1, -1}
	}
	dp[0][code(s[0])] = pair{0, -1}

	for i := 1; i < n; i++ {
		x := code(s[i])
		a, b := -1, -1
		for j := range k {
			if a < 0 || dp[i-1][j].first < dp[i-1][a].first {
				b = a
				a = j
			} else if b < 0 || dp[i-1][j].first < dp[i-1][b].first {
				b = j
			}
		}
		for j := range k {
			var add int
			if j != x {
				add = 1
			}
			if j == a {
				dp[i][j] = pair{dp[i-1][b].first + add, b}
			} else {
				dp[i][j] = pair{dp[i-1][a].first + add, a}
			}
		}
	}

	var ans int

	for i := range k {
		if dp[n-1][i].first < dp[n-1][ans].first {
			ans = i
		}
	}

	best := dp[n-1][ans].first
	buf := []byte(s)

	for i := n - 1; i >= 0; i-- {
		buf[i] = byte(ans + 'A')
		ans = dp[i][ans].second
	}

	return best, string(buf)
}
