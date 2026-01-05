package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
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

func drive(reader *bufio.Reader) int {
	n, _ := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

const inf = 1 << 60

func solve(a []string) int {
	n := len(a)
	m := len(a[0]) - 2

	dp := []int{0, inf}

	get := func(i int) (l int, r int) {
		first := -1
		last := -1
		for j := 1; j <= m; j++ {
			if a[i][j] == '1' {
				if first == -1 {
					first = j
				}
				last = j
			}
		}
		return first, last
	}

	hi := -1
	floors := make([][]int, n)
	for i := range n {
		l, r := get(i)
		floors[i] = []int{l, r}
		if hi == -1 && l != -1 {
			hi = i
		}
	}

	if hi == -1 {
		// no lights on
		return 0
	}

	for i := n - 1; i >= 0; i-- {
		// first and last 1
		l, r := floors[i][0], floors[i][1]
		if hi == i {
			return min(dp[0]+r, dp[1]+m+1-l)
		}
		if l != -1 {
			x := min(dp[0]+m+1, dp[1]+2*(m+1-l))
			y := min(dp[1]+m+1, dp[0]+2*r)
			dp[0] = y
			dp[1] = x
		}
		dp[0]++
		dp[1]++
	}

	return 0
}
