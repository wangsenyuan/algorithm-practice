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
	dp := make([]int, 8)
	ndp := make([]int, 8)
	for i := range dp {
		dp[i] = inf
		ndp[i] = inf
	}
	dp[0] = 0
	dist := make([]int, 3)
	n := len(a[0])
	for _, s := range a {
		for i := range dist {
			dist[i] = inf
		}
		for i := 0; i < n; i++ {
			if s[i] >= '0' && s[i] <= '9' {
				dist[0] = min(dist[0], i, n-i)
			} else if s[i] >= 'a' && s[i] <= 'z' {
				dist[1] = min(dist[1], i, n-i)
			} else {
				dist[2] = min(dist[2], i, n-i)
			}
		}
		for state, v := range dp {
			for i := range 3 {
				if (state>>i)&1 == 0 {
					ndp[state|(1<<i)] = min(ndp[state|(1<<i)], v+dist[i])
				}
			}
		}

		for i := range ndp {
			dp[i] = min(dp[i], ndp[i])
			ndp[i] = inf
		}
	}

	return dp[7]
}
