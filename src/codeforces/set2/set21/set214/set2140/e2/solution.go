package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	var k int
	fmt.Fscan(reader, &k)
	c := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &c[i])
	}
	return solve(n, m, c)
}

const mod = 1e9 + 7

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

func solve(n int, m int, c []int) int {
	N := 1 << n
	f := make([]byte, N)
	f[1] = 1
	for sz := 2; sz <= n; sz++ {
		if (n-sz)%2 == 0 {
		o:
			for mask := 1<<sz - 1; mask > 0; mask-- {
				for _, i := range c {
					i--
					if i >= sz {
						break
					}
					if f[mask>>(i+1)<<i|mask&(1<<i-1)] > 0 {
						f[mask] = 1
						continue o
					}
				}
				f[mask] = 0
			}
		} else {
		o2:
			for mask := 1<<sz - 1; mask > 0; mask-- {
				for _, i := range c {
					i--
					if i >= sz {
						break
					}
					if f[mask>>(i+1)<<i|mask&(1<<i-1)] == 0 {
						f[mask] = 0
						continue o2
					}
				}
				f[mask] = 1
			}
		}
	}

	cnt := make([]int, n+1)
	for s, v := range f {
		cnt[bits.OnesCount(uint(s))] += int(v)
	}

	var res int
	for i, c := range cnt {
		var s int
		for x := 1; x <= m; x++ {
			s = add(s, mul(pow(x-1, n-i), pow(m-x+1, i)))
		}
		res = add(res, mul(s, c))
	}

	return res
}
