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
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

func solve(a []int) int {
	// a[i] <= 20
	n := len(a)
	w := a[0]
	for i := 0; i < n; i++ {
		w = max(w, a[i])
	}
	next := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		next[i] = make([]int, w+1)
		for j := 0; j <= w; j++ {
			next[i][j] = n
		}
	}
	for i := n - 1; i >= 0; i-- {
		x := a[i] - 1
		copy(next[i], next[i+1])
		next[i][x] = i
	}
	W := 1 << w

	dp := make([]int, W)
	for i := 0; i < W; i++ {
		dp[i] = n
	}
	dp[0] = -1

	for state := 0; state < W; state++ {
		for i := 0; i < w; i++ {
			if (state>>i)&1 == 1 {
				continue
			}
			j := next[dp[state]+1][i]
			j = next[j+1][i]
			next := state | 1<<i
			dp[next] = min(dp[next], j)
		}
	}

	var res int

	for state := 1; state < W; state++ {
		if dp[state] < n {
			tmp := state
			var cnt int
			for tmp > 0 {
				cnt++
				tmp -= tmp & -tmp
			}
			res = max(res, 2*cnt)
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
