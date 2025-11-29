package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	S := make([]string, n)
	for i := 0; i < n; i++ {
		S[i] = readString(reader)
	}
	res := solve(S)
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

const MOD = 998244353

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func solve(S []string) int64 {
	n := len(S)
	cnt := make([][26]int, n)
	for i, s := range S {
		for j := range len(s) {
			cnt[i][int(s[j]-'a')]++
		}
	}

	f := make([]int, 1<<n)
	mn := [26]int{}

	for mask := 1; mask < 1<<n; mask++ {
		for i := range 26 {
			mn[i] = INF
		}
		for m := uint(mask); m > 0; m &= m - 1 {
			t := bits.TrailingZeros(m)
			for i, c := range cnt[t][:] {
				mn[i] = min(mn[i], c)
			}
		}

		res := bits.OnesCount(uint(mask))%2*2 - 1 + MOD
		for i := range 26 {
			res = modMul(res, mn[i]+1)
		}
		f[mask] = res
	}

	for i := range n {
		for s := 0; s < 1<<n; s++ {
			s |= 1 << i
			f[s] = modAdd(f[s], f[s^(1<<i)])
		}
	}
	var ans int

	for mask, v := range f {
		var sum int
		for m := uint(mask); m > 0; m &= m - 1 {
			sum += bits.TrailingZeros(m)
		}
		k := bits.OnesCount(uint(mask))
		ans ^= v * k * (sum + k)
	}

	return int64(ans)
}

func solve1(S []string) int64 {
	n := len(S)
	// n <= 23
	// f([t1....tm]) = number of strings, being sub-seq of at least one of ti
	// g([t1...tm]) = number of strings, being sub-seq of all of ti
	// f = state over m inclusion-exclusion such things
	// how to get g?
	// len(S[i]) <= 10000, S[i] is sorted
	P := make([][]int, n)
	for i := 0; i < n; i++ {
		P[i] = convert(S[i])
	}
	// g = prod[0..26] pow(2, min(cnt[i])) - 1
	T := 1 << n
	H := make([]int, T)
	cnt := make([]int, 26)
	for i := 0; i < 26; i++ {
		cnt[i] = INF
	}
	for state := 1; state < T; state++ {

		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				for j := 0; j < 26; j++ {
					cnt[j] = min(cnt[j], P[i][j])
				}
			}
		}
		H[state] = 1
		for i := 0; i < 26; i++ {
			H[state] = modMul(H[state], cnt[i]+1)
			cnt[i] = INF
		}
	}

	flipAll(H, n)

	for i := n - 1; i >= 0; i-- {
		for j := (1 << n) - 1; j >= 0; j-- {
			if (j>>i)&1 == 1 {
				H[j] = modSub(H[j], H[j^(1<<i)])
			}
		}
	}

	flipAll(H, n)
	var sum int

	for i := 0; i < T; i++ {
		sum = modAdd(sum, H[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 1<<n; j++ {
			if (j>>i)&1 == 0 {
				H[j^(1<<i)] = modAdd(H[j^(1<<i)], H[j])
			}
		}
	}

	res := make([]int, T)

	for i := 0; i < T; i++ {
		res[i] = modSub(sum, H[(T-1)^i])
	}

	var ans int64

	for i := 0; i < T; i++ {
		var c, s int
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				c++
				s += j + 1
			}
		}
		ans ^= int64(res[i]) * int64(c) * int64(s)
	}

	return ans
}

func flipAll(dp []int, n int) {
	mask := (1 << n) - 1
	for i := 0; i < (1 << (n - 1)); i++ {
		j := i ^ mask
		dp[i], dp[j] = dp[j], dp[i]
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const INF = 1 << 30

func convert(s string) []int {
	res := make([]int, 26)

	for i, j := 0, 0; i < 26 && j < len(s); i++ {
		for j < len(s) && int(s[j]-'a') == i {
			j++
			res[i]++
		}
	}
	return res
}
