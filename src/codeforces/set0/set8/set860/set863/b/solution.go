package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	w := readNNums(reader, 2*n)
	return solve(w)
}

func solve(w []int) int {
	sort.Ints(w)
	n := len(w)
	// n 是偶数
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		ndp := make([]int, 3)
		if i == 0 {
			ndp[0] = inf
			ndp[1] = 0
			ndp[2] = inf
		} else if i == 1 {
			ndp[0] = w[i] - w[i-1]
			ndp[1] = inf
			ndp[2] = 0
		} else {
			// i >= 2
			ndp[0] = dp[i-2][0] + w[i] - w[i-1]
			ndp[1] = min(w[i]-w[i-1]+dp[i-2][1], dp[i-1][0])
			ndp[2] = min(w[i]-w[i-1]+dp[i-2][2], dp[i-1][1])
		}
		dp[i] = ndp
	}

	return dp[n-1][2]
}

const inf = 1 << 30
