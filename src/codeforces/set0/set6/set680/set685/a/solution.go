package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n, &m)
	res := solve(n, m)
	fmt.Println(res)
}

func solve(n int, m int) int {
	dp := calc(n)
	fp := calc(m)

	var res int

	for s1 := range 1 << 7 {
		for s2 := range 1 << 7 {
			if s1*s2 > 0 && s1&s2 == 0 {
				res += dp[s1] * fp[s2]
			}
		}
	}

	return res
}

func calc(n int) []int {
	// 0到n-1
	n--
	var ds []int
	for i := n; i > 0; i /= 7 {
		ds = append(ds, i%7)
	}
	if n == 0 {
		ds = append(ds, 0)
	}

	slices.Reverse(ds)

	dp := make([]int, (1<<7)*2)
	dp[0] = 1

	ndp := make([]int, (1<<7)*2)

	for i := 0; i < len(ds); i++ {
		for lt := range 2 {
			for mask := range 1 << 7 {
				if dp[mask*2+lt] == 0 {
					continue
				}
				end := 7
				if lt == 0 {
					end = ds[i] + 1
				}
				for v := range end {
					if (mask>>v)&1 == 0 {
						newMask := mask | (1 << v)
						nlt := lt
						if v < ds[i] {
							nlt = 1
						}
						w := newMask*2 + nlt
						ndp[w] += dp[mask*2+lt]
					}
				}
			}
		}
		copy(dp, ndp)
		clear(ndp)
	}

	res := make([]int, 1<<7)

	for mask := range 1 << 7 {
		// <= n都是可以的
		res[mask] = dp[mask*2] + dp[mask*2+1]
	}

	return res
}
