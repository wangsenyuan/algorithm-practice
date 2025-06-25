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
	m := readNum(reader)
	X := readNNums(reader, m)
	Y := readNNums(reader, m)
	return solve(X, Y)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(X []int, Y []int) int {
	// m := len(X)
	var n int
	for _, x := range X {
		n += x
	}

	C := make([][]int, n+1)
	for i := range n + 1 {
		C[i] = make([]int, i+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = add(C[i-1][j-1], C[i-1][j])
		}
	}

	dp := make([]int, n+1)
	dp[0] = 1
	ndp := make([]int, n+1)
	sum := n
	f := 1
	for i, x := range X {
		f = mul(f, C[sum][x])
		sum -= x

		clear(ndp)
		for w := 0; w+x <= n; w++ {
			nw := w + x
			for u := 0; u <= Y[i] && nw-u >= 0; u++ {
				// 选择其中u个人参加第二堂课
				ndp[nw-u] = add(ndp[nw-u], mul(dp[w], C[nw][u]))
			}
		}
		copy(dp, ndp)
	}
	return mul(dp[0], f)
}
