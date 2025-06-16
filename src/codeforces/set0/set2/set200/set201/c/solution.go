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
	a := readNNums(reader, n-1)
	return solve(n, a)
}

func solve(n int, a []int) int {
	fp := make([][2]int, n)
	for i := 1; i < n-1; i++ {
		if a[i-1] == 1 {
			fp[i][0] = 0
			// fp[i][1]表示如果从i出发，但是不用返回时的最大值
			fp[i][1] = 1 + max(fp[i-1][0], fp[i-1][1])
		} else {
			fp[i][0] = a[i-1]/2*2 + fp[i-1][0]
			fp[i][1] = a[i-1] - (1 - a[i-1]&1) + max(fp[i-1][0], fp[i-1][1])
		}
	}
	ans := max(fp[n-2][0], fp[n-2][1])

	dp := []int{0, 0}

	for i := n - 2; i >= 0; i-- {
		if a[i] == 1 {
			dp[1] = 1 + max(dp[0], dp[1])
			dp[0] = 0
		} else {
			dp[1] = a[i] - (1 - a[i]&1) + max(dp[0], dp[1])
			dp[0] += a[i] / 2 * 2
		}
		if i > 0 {
			// 如果从i开始，先往前的话，必须回来，再往后，反之依然
			ans = max(ans, fp[i][0]+dp[1], fp[i][1]+dp[0])
		} else {
			ans = max(ans, dp[0], dp[1])
		}
	}

	return ans
}
