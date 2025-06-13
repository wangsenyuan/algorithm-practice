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

const mod = 998244353

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
	n, m := readTwoNums(reader)
	c := readNNums(reader, m)
	return solve(n, m, c)
}

func solve(n int, m int, c []int) int {
	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, n+1)
		for j := range n + 1 {
			dp[i][j] = -1
		}
	}

	pos := make([]int, n+1)
	for i := range c {
		pos[c[i]] = i
	}

	var dfs func(l int, r int) int
	dfs = func(l int, r int) (ans int) {
		if l >= r {
			return 1
		}
		if dp[l][r] != -1 {
			return dp[l][r]
		}
		defer func() {
			dp[l][r] = ans
		}()

		x := m
		for i := l; i <= r; i++ {
			x = min(x, c[i])
		}

		var tmp1 int
		for a := l; a <= pos[x]; a++ {
			tmp1 = add(tmp1, mul(dfs(l, a-1), dfs(a, pos[x]-1)))
		}
		var tmp2 int
		for b := pos[x]; b <= r; b++ {
			tmp2 = add(tmp2, mul(dfs(pos[x]+1, b), dfs(b+1, r)))
		}
		// dp[l][a] * dp[a+1][pos[x] - 1] * dp[pos[x] + 1][b] * dp[b+1][r]
		ans = mul(tmp1, tmp2)
		return
	}

	return dfs(0, n-1)
}
