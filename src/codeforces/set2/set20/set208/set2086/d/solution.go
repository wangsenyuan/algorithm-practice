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
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	buf.WriteTo(os.Stdout)
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
	c := readNNums(reader, 26)
	return solve(c)
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
	return int(int64(a) * int64(b) % mod)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func inv(n int) int {
	return pow(n, mod-2)
}

const N = 500010

var F [N]int
var I [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(i, F[i-1])
	}
	I[N-1] = pow(F[N-1], mod-2)
	for i := N - 2; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}
}

func nCr(n int, r int) int {
	if n < r || r < 0 {
		return 0
	}

	return mul(F[n], mul(I[r], I[n-r]))
}

func solve(c []int) int {
	var n int
	for _, v := range c {
		n += v
	}
	h := n / 2

	dp := make([]int, h+1)
	dp[0] = 1
	for _, v := range c {
		if v > 0 {
			for j := h; j >= v; j-- {
				dp[j] = add(dp[j], dp[j-v])
			}
		}
	}
	ans := mul(dp[h], mul(F[h], F[n-h]))
	for _, v := range c {
		ans = mul(ans, inv(F[v]))
	}

	return ans
}

func solve1(c []int) int {
	var n int
	for _, v := range c {
		n += v
	}
	h := n / 2

	// dp[diff] = 到目前为止, 奇数为止 - 偶数位置 = diff 时的计数
	dp := make([]int, 2*n+1)
	dp[n] = 1

	fp := make([]int, 2*n+1)
	var sum int
	for _, v := range c {
		if v == 0 {
			continue
		}
		clear(fp)
		sum += v
		// a + b = sum
		// a - b = diff
		// 把当前的v放置在奇数为止
		for diff := -sum; diff <= sum; diff++ {
			a := (sum + diff) / 2
			if a >= v && a <= h {
				fp[n+diff] = add(fp[n+diff], mul(dp[n+diff-v], nCr(a, v)))
			}
			b := sum - a
			if b >= v && b <= n-h {
				fp[n+diff] = add(fp[n+diff], mul(dp[diff+n+v], nCr(b, v)))
			}
		}

		copy(dp, fp)
	}

	if 2*h == n {
		return dp[n]
	}
	return dp[n-1]
}
