package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	os.Stdout.Write(buf.Bytes())
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n-1)
	return solve(n, a)
}

func solve(n int, a []int) []int {
	arr := make([]int, n+1)
	copy(arr[2:], a)
	a = arr
	dp := make([][]int, n+1)
	vis := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, 2)
		vis[i] = make([]int, 2)
	}

	// dfs(u) 返回x = u时，y的增加值
	var dfs func(u int, d int) int

	dfs = func(u int, d int) int {
		if u == 1 {
			// 如果回到位置1，如果要重新往后移动，那么就陷入了循环
			// 否则就是往前移动，肯定移处了区间
			if d == 0 {
				return -inf
			}
			return arr[u]
		}
		if vis[u][d] == 1 {
			// 陷入了循环
			return -inf
		}
		if vis[u][d] == 2 {
			// u/d有可能是个循环
			return dp[u][d]
		}
		vis[u][d]++
		// y增加 arr[u]
		dp[u][d] = arr[u]
		var v int
		// dp[u] = -1
		if d == 0 {
			v = u + arr[u]
		} else {
			v = u - arr[u]
		}
		if v >= 1 && v <= n {
			dp[u][d] += dfs(v, d^1)
		}
		vis[u][d]++
		return dp[u][d]
	}

	ans := make([]int, n)

	for i := 1; i < n; i++ {
		arr[1] = i
		ans[i] = max(-1, i+dfs(i+1, 1))
	}

	return ans[1:]
}

const inf = 1 << 60
