package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for range tc {
		x, y := readTwoNums(reader)
		res := solve(x, y)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

const H = 60

const inf = 1 << 60

var dp [][]int

func init() {

	dp = make([][]int, H)
	for i := range H {
		dp[i] = make([]int, H)
		for j := range H {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0

	for i := 0; i < H; i++ {
		for j := 0; j < H; j++ {
			// dp[i][j]已知
			v := dp[i][j]
			if v == inf {
				continue
			}
			for x := 1; x < H; x++ {
				if (v>>x)&1 == 0 {
					// x还可以使用
					if i+x < H {
						dp[i+x][j] = min(dp[i+x][j], v|(1<<x))
					}
					if j+x < H {
						dp[i][j+x] = min(dp[i][j+x], v|(1<<x))
					}
				}
			}
		}
	}
}

func solve(x int, y int) int {
	// 假设对x处理了多次，分别是 k1, k2, k3...
	// 总体的效果相当于 x >> (k1 + k2 + ...)
	ans := inf
	for k1 := 0; k1 < H; k1++ {
		for k2 := 0; k2 < H; k2++ {
			if (x >> k1) == (y >> k2) {
				ans = min(ans, dp[k1][k2])
			}
		}
	}

	return ans
}
