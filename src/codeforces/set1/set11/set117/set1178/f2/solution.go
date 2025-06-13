package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	c = slices.Compact(c)
	m = len(c)
	if m > 2*n {
		return 0
	}
	pos := make([][]int, n)
	for i := range m {
		c[i]--
		pos[c[i]] = append(pos[c[i]], i)
	}
	// c[i] < n
	dp := make([][]int, m+1)
	for i := range m + 1 {
		dp[i] = make([]int, m+1)
		for j := range m + 1 {
			dp[i][j] = 1
		}
	}

	for l := 1; l <= m; l++ {
		for a := 0; a+l <= m; a++ {
			lo := c[a]
			for i := 0; i < l; i++ {
				lo = min(lo, c[a+i])
			}
			j := pos[lo][0] - a
			k := pos[lo][len(pos[lo])-1] - a
			if j < 0 || k >= l {
				dp[a][l] = 0
				continue
			}
			var left, right int
			for u := 0; u <= j; u++ {
				left = add(left, mul(dp[a][u], dp[a+u][j-u]))
			}
			for v := k + 1; v <= l; v++ {
				right = add(right, mul(dp[a+k+1][v-(k+1)], dp[a+v][l-v]))
			}

			dp[a][l] = mul(left, right)
			for w := 0; w < len(pos[lo])-1; w++ {
				if pos[lo][w]+1 != pos[lo][w+1] {
					tmp := dp[pos[lo][w]+1][pos[lo][w+1]-pos[lo][w]-1]
					dp[a][l] = mul(dp[a][l], tmp)
				}
			}
		}
	}

	return dp[0][m]
}

func sortAndUnique(c []int) []int {
	res := slices.Clone(c)
	slices.Sort(res)
	return slices.Compact(res)
}
