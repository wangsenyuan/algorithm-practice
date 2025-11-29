package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var l, r, k int
	fmt.Fscan(reader, &l, &r, &k)
	return solve(l, r, k)
}

const mod = 998244353

func add(nums ...int) int {
	var res int
	for _, num := range nums {
		res += num
		res %= mod
	}
	return res
}
func mul(nums ...int) int {
	res := 1
	for _, num := range nums {
		res *= num
		res %= mod
	}
	return res
}

func solve(l int, r int, k int) int {

	pw := make([]int, 20)

	pw[0] = 1
	for i := 1; i < 20; i++ {
		pw[i] = mul(pw[i-1], 10)
	}

	calc := func(r int) int {
		var dp [1 << 10][2][2]int
		var ndp [1 << 10][2][2]int
		var fp [1 << 10][2][2]int
		var nfp [1 << 10][2][2]int

		dp[0][1][1] = 1
		fp[0][1][1] = 0

		var ds []int
		for i := r; i > 0; i /= 10 {
			ds = append(ds, i%10)
		}

		slices.Reverse(ds)
		n := len(ds)

		for iter, d := range ds {
			for s := range 1 << 10 {
				for lz := range 2 {
					for eq := range 2 {
						lo, hi := 0, 9
						if eq == 1 {
							hi = d
						}
						for d1 := lo; d1 <= hi; d1++ {
							var nlz int
							if lz == 1 && d1 == 0 {
								nlz = 1
							}
							neq := eq
							if d1 < d {
								neq = 0
							}
							ns := s
							if lz == 0 || d1 > 0 {
								ns |= 1 << d1
							}
							ndp[ns][nlz][neq] = add(ndp[ns][nlz][neq], dp[s][lz][eq])
							nfp[ns][nlz][neq] = add(nfp[ns][nlz][neq], fp[s][lz][eq], mul(dp[s][lz][eq], d1, pw[n-iter-1]))
						}
					}
				}
			}
			for s := range 1 << 10 {
				for lz := range 2 {
					for eq := range 2 {
						dp[s][lz][eq] = ndp[s][lz][eq]
						ndp[s][lz][eq] = 0
						fp[s][lz][eq] = nfp[s][lz][eq]
						nfp[s][lz][eq] = 0
					}
				}
			}
		}

		var ans int
		for s := range 1 << 10 {
			if bits.OnesCount(uint(s)) > k {
				continue
			}
			for eq := range 2 {
				ans = add(ans, fp[s][0][eq])
			}
		}
		return ans
	}

	return add(calc(r), mod-calc(l-1))
}
