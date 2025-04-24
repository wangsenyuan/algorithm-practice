package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Printf("%.10f\n", res)
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

func process(reader *bufio.Reader) float64 {
	n, D := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(a, D)
}

const inf = 1e16

func solve(a []int, D int) float64 {
	n := len(a)
	N := 1 << n
	dp := make([][]float64, N)
	for i := range dp {
		dp[i] = make([]float64, D+1)
	}
	var sum int
	for _, v := range a {
		sum += v
	}
	avg := float64(sum) / float64(D)

	dp[0][0] = 0

	for mask := 0; mask < N; mask++ {
		var tmp int
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 {
				tmp += a[i]
			}
		}
		dp[mask][1] = math.Pow(float64(tmp)-avg, 2)
		for d := 2; d <= D; d++ {
			dp[mask][d] = dp[mask][d-1] + dp[0][1]
			for sub := mask; sub > 0; sub = (sub - 1) & mask {
				dp[mask][d] = math.Min(dp[mask][d], dp[sub][d-1]+dp[mask^sub][1])
			}
		}
	}

	return dp[N-1][D] / float64(D)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
