package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	board := make([]string, 2)
	for i := 0; i < 2; i++ {
		board[i] = readString(reader)[:n]
	}
	res := solve(board)
	fmt.Println(res)
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(board []string) int {
	n := len(board[0])
	f := make([][2]int, n+1)
	for j := n - 2; j >= 0; j-- {
		for i := range 2 {
			f[j][i] = f[j+1][i] + int(board[i^1][j]&1)
			if board[i^1][j] == '1' {
				f[j][i] = min(f[j][i], f[j+2][i^1]+int(board[i][j+1]&1))
			}
		}
	}
	return strings.Count(board[0], "1") + strings.Count(board[1], "1") - f[0][0]
}

func solve1(board []string) int {
	n := len(board[0])
	if n == 1 {
		return 0
	}

	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, 2)
			for k := 0; k < 2; k++ {
				dp[i][j][k] = -n
			}
		}
	}
	if board[1][0] == '1' {
		dp[0][0][1] = 1
	}
	dp[0][0][0] = 0

	getValue := func(i, j int) int {
		return int(board[i][j] - '0')
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < 2; j++ {
			j0 := getValue(j, i+1)
			j1 := getValue(1^j, i+1)
			dp[i+1][j^1][0] = max(dp[i+1][j^1][0], dp[i][j][1]+j1)
			dp[i+1][j][j1] = max(dp[i+1][j][j1], dp[i][j][0]+j1+j0)
			dp[i+1][j][0] = max(dp[i+1][j][0], dp[i][j][0]+j0)
		}
	}

	return max(dp[n-1][0][0], dp[n-1][0][1], dp[n-1][1][0], dp[n-1][1][1])
}

func max(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
