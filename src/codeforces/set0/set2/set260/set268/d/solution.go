package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, h := readTwoNums(reader)
	res := solve(n, h)
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

const mod = 1e9 + 9

func add(a, b int) int {
	return (a + b) % mod
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(nums ...int) int {
	res := 1
	for _, v := range nums {
		res *= v
		res %= mod
	}
	return res
}

func solve(n int, h int) int {
	dp := make([][][][]int, 2)
	ndp := make([][][][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][][]int, h+1)
		ndp[i] = make([][][]int, h+1)
		for j := range h + 1 {
			dp[i][j] = make([][]int, h+1)
			ndp[i][j] = make([][]int, h+1)
			for k := range h + 1 {
				dp[i][j][k] = make([]int, h+1)
				ndp[i][j][k] = make([]int, h+1)
			}
		}
	}
	// first is for alive or not
	dp[1][0][0][0] = 4

	inc := func(j int) int {
		if j == 0 {
			return 0
		}
		return (j + 1) % h
	}

	use := func(i int, j int) int {
		if j == 0 && i < h || j > 0 {
			return 1
		}
		return 0
	}

	next := make([]int, h)
	for i := 0; i < h; i++ {
		next[i] = inc(i)
	}

	to := make([][]int, n)
	for i := range n {
		to[i] = make([]int, h)
		for j := range h {
			to[i][j] = use(i, j)
		}
	}

	for i := 1; i < n; i++ {
		for alive := range 2 {
			for h1 := range h {
				for h2 := range h {
					for h3 := range h {
						nh1 := next[h1]
						nh2 := next[h2]
						nh3 := next[h3]
						ndp[alive][nh1][nh2][nh3] = add(ndp[alive][nh1][nh2][nh3], dp[alive][h1][h2][h3])
						// 使用h1
						na1 := to[i][h1]
						h4 := 1 * alive
						ndp[na1][h4][nh2][nh3] = add(ndp[na1][h4][nh2][nh3], dp[alive][h1][h2][h3])
						na2 := to[i][h2]
						ndp[na2][nh1][h4][nh3] = add(ndp[na2][nh1][h4][nh3], dp[alive][h1][h2][h3])
						na3 := to[i][h3]
						ndp[na3][nh1][nh2][h4] = add(ndp[na3][nh1][nh2][h4], dp[alive][h1][h2][h3])
					}
				}
			}
		}
		for alive := range 2 {
			for h1 := range h {
				for h2 := range h {
					for h3 := range h {
						dp[alive][h1][h2][h3] = ndp[alive][h1][h2][h3]
						ndp[alive][h1][h2][h3] = 0
					}
				}
			}
		}
	}

	var ans int

	for h1 := range h {
		for h2 := range h {
			for h3 := range h {
				ans = add(ans, dp[1][h1][h2][h3])
				if h1+h2+h3 > 0 {
					ans = add(ans, dp[0][h1][h2][h3])
				}
			}
		}
	}

	return ans
}
