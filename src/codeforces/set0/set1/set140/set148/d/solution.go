package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	w, b := readTwoNums(reader)
	res := solve(w, b)
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

func solve(w int, b int) float64 {
	dp := make([][]float64, w+1)
	for i := range w + 1 {
		dp[i] = make([]float64, b+1)
		for j := range b + 1 {
			dp[i][j] = -1
		}
	}

	var dfs func(x int, y int) float64

	dfs = func(x int, y int) (ans float64) {
		if x <= 0 {
			return 0
		}
		if y <= 0 {
			return 1.0
		}

		if dp[x][y] >= 0 {
			return dp[x][y]
		}

		defer func() {
			dp[x][y] = ans
		}()
		// 如果公主抽到了白色
		ans = float64(x) / float64(x+y)
		// y > 1
		// 公主抽到黑色的概率
		lose := float64(y) / float64(x+y)
		// 剩下(x, y - 1)
		// 且龙抽走的也是黑色
		lose *= float64(y-1) / float64(x+y-1)

		if lose >= 1e-13 {
			// 剩下(x, y - 2)
			tmp := dfs(x-1, y-2) * float64(x) / float64(x+y-2)
			tmp += dfs(x, y-3) * float64(y-2) / float64(x+y-2)
			ans += lose * tmp
		}

		return ans
	}

	return dfs(w, b)
}
