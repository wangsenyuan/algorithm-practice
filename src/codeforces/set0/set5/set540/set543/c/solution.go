package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
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
	n, m := readTwoNums(reader)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = readString(reader)
	}
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = readNNums(reader, m)
	}
	return solve(s, a)
}

const inf = 1 << 60

func solve(s []string, a [][]int) int {
	n := len(s)
	m := len(s[0])

	type pair struct {
		first  int
		second int
	}

	col := make([][]pair, n)
	for i := range n {
		col[i] = make([]pair, m)
		for j := range m {
			x := s[i][j]
			var sum int
			var mx int
			var mask int
			for i1 := range n {
				if s[i1][j] == x {
					sum += a[i1][j]
					mx = max(mx, a[i1][j])
					mask |= 1 << i1
				}
			}
			col[i][j] = pair{sum - mx, mask}
		}
	}
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	for state := 0; state < (1<<n)-1; state++ {
		i := bits.TrailingZeros(^uint(state))
		dp[state|1<<i] = min(dp[state|1<<i], dp[state]+slices.Min(a[i]))
		for _, c := range col[i] {
			dp[state|c.second] = min(dp[state|c.second], dp[state]+c.first)
		}
	}
	return dp[(1<<n)-1]
}
