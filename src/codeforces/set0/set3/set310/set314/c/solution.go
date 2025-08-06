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

func process(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, a)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
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

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] = add(bit[i], v)
		i += i & -i
	}
}

func (bit BIT) query(i int) int {
	i++
	ans := 0
	for i > 0 {
		ans = add(ans, bit[i])
		i -= i & -i
	}
	return ans
}

func solve(n int, a []int) int {
	x := slices.Max(a)
	// dp[x]表示最后一个v出现的序列的数量
	fp := make(BIT, x+2)
	for _, v := range a {
		// 加上当前自己
		old := sub(fp.query(v), fp.query(v-1))

		ways := fp.query(v) + 1
		// 为了避免重复，把之前v的计数给减去

		// 这里还需要知道之前的ways
		ways = mul(v, ways)
		// fp.update(v, ways)
		fp.update(v, sub(ways, old))
	}

	return fp.query(x)
}
