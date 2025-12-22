package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
}

func process(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(m, a)
}

const inf = 1 << 60

func solve(m int, a []int) int {
	n := len(a)
	sum := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		sum[i] = make([]int, m)
	}
	for i, x := range a {
		copy(sum[i+1], sum[i])
		sum[i+1][x-1]++
	}

	u := 1 << m
	sz := make([]int, u)
	for i, v := range sum[n][:m] {
		// v是对应i的数量
		highBit := 1 << i
		for mask, s := range sz[:highBit] {
			sz[highBit|mask] = s + v
		}
	}

	f := make([]int, u)
	for s, fs := range f {
		for cus, lb := u-1^s, 0; cus > 0; cus ^= lb {
			lb = cus & -cus
			ns := s | lb
			p := bits.TrailingZeros(uint(lb))
			f[ns] = max(f[ns], fs+sum[sz[ns]][p]-sum[sz[s]][p])
		}
	}

	return n - f[u-1]
}
