package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n, x int
	fmt.Fscan(reader, &n, &x)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, x)
}

const mod = 1_000_000_007

const N = 500010

var lpf [N]int
var primes []int

func init() {
	for i := 2; i < N; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= N {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}
}

func mul(a, b int) int {
	return a * b % mod
}

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(a []int, x int) int {
	type primePower struct {
		prime int
		exp   int
	}

	var factors []primePower
	x0 := x
	for x0 > 1 {
		p := lpf[x0]
		var c int
		for x0%p == 0 {
			c++
			x0 /= p
		}
		factors = append(factors, primePower{p, c})
	}

	pos := make(map[int]int, len(factors))
	for i, cur := range factors {
		pos[cur.prime] = i
	}

	f0 := make(map[int]int)
	cnt := make([][20]int, len(factors))
	mxExp := make([]int, len(factors))

	for _, v := range a {
		for v > 1 {
			p := lpf[v]
			var c int
			for v%p == 0 {
				c++
				v /= p
			}
			f0[p] += c
			if j, ok := pos[p]; ok {
				cnt[j][c]++
				mxExp[j] = max(mxExp[j], c)
			}
		}
	}

	mulPoly := func(a, b []int) []int {
		c := make([]int, len(a))
		for i, x := range a {
			if x == 0 {
				continue
			}
			for j, y := range b[:len(b)-i] {
				if y != 0 {
					c[i+j] = add(c[i+j], mul(x, y))
				}
			}
		}
		return c
	}

	powPoly := func(base []int, exp int) []int {
		res := make([]int, len(base))
		res[0] = 1
		for exp > 0 {
			if exp&1 == 1 {
				res = mulPoly(res, base)
			}
			exp >>= 1
			if exp > 0 {
				base = mulPoly(base, base)
			}
		}
		return res
	}

	play := func(id int, up int) int {
		var ans int
		var prev []int
		lim := mxExp[id] + up
		for mx := 0; mx <= mxExp[id]; mx++ {
			dp := make([]int, lim+1)
			dp[0] = 1
			for e, c := range cnt[id] {
				cap := min(e, mx)
				if c == 0 || cap == 0 {
					continue
				}
				base := make([]int, lim+1)
				for i := 0; i <= cap && i <= lim; i++ {
					base[i] = 1
				}
				dp = mulPoly(dp, powPoly(base, c))
			}
			target := mx + up
			cur := dp[target]
			if prev != nil {
				cur = add(cur, mod-prev[target])
			}
			ans = add(ans, cur)
			prev = dp
		}
		return ans
	}

	ans := 1
	for k, v := range f0 {
		if _, ok := pos[k]; !ok {
			ans = mul(ans, v+1)
		}
	}
	for i, cur := range factors {
		ans = mul(ans, play(i, cur.exp))
	}

	return ans
}
