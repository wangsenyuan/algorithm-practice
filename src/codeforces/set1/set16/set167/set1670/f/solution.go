package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line := readNInt64s(reader, 4)
	res := solve(int(line[0]), line[1], line[2], line[3])
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(n int, l int64, r int64, z int64) int {
	ans := compute(n, int(r), int(z))
	ans = modAdd(ans, MOD-compute(n, int(l-1), int(z)))
	return ans
}

func compute(n int, r int, z int) int {
	if r < z {
		return 0
	}
	m := bits.Len(uint(r))
	dp := make([][][2]int, m)
	for i := range dp {
		dp[i] = make([][2]int, n)
		for j := range dp[i] {
			dp[i][j] = [2]int{-1, -1}
		}
	}
	var f func(int, int, int) int
	f = func(i, s, gr int) (res int) {
		if i == m {
			if s == 0 && gr == 0 {
				return 1
			}
			return 0
		}
		p := &dp[i][s][gr]
		if *p >= 0 {
			return *p
		}
		for j := z >> i & 1; j <= n; j += 2 {
			bit := (s + j) & 1
			newGr := 0
			if bit == r>>i&1 {
				newGr = gr
			} else if bit > r>>i&1 {
				newGr = 1
			}
			res = (res + f(i+1, (s+j)>>1, newGr)*nCr(n, j)) % MOD
		}
		*p = res
		return
	}
	return f(0, 0, 0)
}

func compute1(n int, v int64, z int64) int {
	dp := make([][]int, H)
	for i := 0; i < H; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(id int, rem int) int

	dfs = func(id int, rem int) int {
		if rem < 0 {
			return 0
		}
		if id < 0 {
			return 1
		}
		if rem >= N {
			rem = N - 1
		}
		if dp[id][rem] < 0 {
			var ans int
			curBitXor := int((z >> id) & 1)
			for i := curBitXor; i <= n; i += 2 {
				curBitSum := int((v >> id) & 1)
				nextRem := 2 * (rem + curBitSum - i)
				add := modMul(nCr(n, i), dfs(id-1, nextRem))
				ans = modAdd(ans, add)
			}
			dp[id][rem] = ans
		}

		return dp[id][rem]
	}

	return dfs(H-1, 0)
}

const H = 61

const N = 2010
const MOD = 1e9 + 7

var F []int
var I []int

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = modMul(res, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return res
}

func init() {
	F = make([]int, N)
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = modMul(i, F[i-1])
	}

	I = make([]int, N)
	I[N-1] = pow(F[N-1], MOD-2)

	for i := N - 2; i >= 0; i-- {
		I[i] = modMul(i+1, I[i+1])
	}
}

func nCr(n, r int) int {
	if n < r {
		return 0
	}
	res := F[n]
	res = modMul(res, I[r])
	res = modMul(res, I[n-r])
	return res
}
