package main

import (
	"bufio"
	"fmt"
	"os"
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

func process(reader *bufio.Reader) int {
	n, k, d := readThreeNums(reader)
	return solve(n, k, d)
}

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(n int, k int, d int) int {
	dp := make([][]int, n+1)
	ndp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, 2)
		ndp[i] = make([]int, 2)
	}
	dp[0][0] = 1
	var res int
	for range n {
		for s := range n {
			for u := 1; u <= k && s+u <= n; u++ {
				if u < d {
					ndp[s+u][0] = add(ndp[s+u][0], dp[s][0])
					ndp[s+u][1] = add(ndp[s+u][1], dp[s][1])
				} else {
					ndp[s+u][1] = add(ndp[s+u][1], add(dp[s][0], dp[s][1]))
				}
			}
		}
		for s := range n + 1 {
			copy(dp[s], ndp[s])
			clear(ndp[s])
		}
		res = add(res, dp[n][1])
	}

	return res
}
