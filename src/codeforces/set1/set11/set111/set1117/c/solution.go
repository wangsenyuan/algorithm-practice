package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	pos1 := readNNums(reader, 2)
	pos2 := readNNums(reader, 2)
	readNum(reader)
	s := readString(reader)
	return solve(pos1, pos2, s)
}

const inf = 1 << 60

func solve(pos1 []int, pos2 []int, s string) int {
	n := len(s)

	dp := make([][]int, n+1)
	dp[0] = make([]int, 4)

	for i := range n {
		dp[i+1] = slices.Clone(dp[i])
		switch s[i] {
		case 'U':
			dp[i+1][0]++
		case 'D':
			dp[i+1][1]++
		case 'L':
			dp[i+1][2]++
		case 'R':
			dp[i+1][3]++
		}
	}

	var dd = [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}

	check := func(m int) bool {
		x, y := pos1[0], pos1[1]
		for i := range 4 {
			dx, dy := dd[i][0], dd[i][1]
			x += m / n * dx * dp[n][i]
			y += m / n * dy * dp[n][i]
		}
		for i := range 4 {
			dx, dy := dd[i][0], dd[i][1]
			x += dx * dp[m%n][i]
			y += dy * dp[m%n][i]
		}

		dx := abs(pos2[0] - x)
		dy := abs(pos2[1] - y)
		return dx+dy <= m
	}

	if !check(inf) {
		return -1
	}
	return sort.Search(inf, check)
}

func abs(num int) int {
	return max(num, -num)
}
