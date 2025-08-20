package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var tc int
	fmt.Fscan(reader, &tc)

	var buf bytes.Buffer

	for range tc {
		var b, l, r int
		fmt.Fscan(reader, &b, &l, &r)
		buf.WriteString(fmt.Sprintf("%d\n", solve(b, l, r)))
	}

	buf.WriteTo(os.Stdout)
}

var fp [][][][]int

func init() {
	fp = make([][][][]int, 11)
	for b := 2; b <= 10; b++ {
		fp[b] = prepare(b)
	}
}

const MAX_NUM = 1e18 + 10

func prepare(b int) [][][]int {
	ml := 1
	for i := b; i < MAX_NUM; i *= b {
		ml++
	}

	dp := make([][][]int, 2)
	for i := range 2 {
		dp[i] = make([][]int, 1<<b)
		for j := range dp[i] {
			dp[i][j] = make([]int, ml+2)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var f func(in int, m int, pos int) int
	f = func(in int, m int, pos int) (res int) {
		if dp[in][m][pos] != -1 {
			return dp[in][m][pos]
		}
		defer func() {
			dp[in][m][pos] = res
		}()

		if pos == 0 {
			// 有非0的数，且所有的数，出现了偶数次
			if in != 0 && m == 0 {
				res = 1
			}
			return
		}

		if in == 0 {
			res += f(0, m, pos-1)
		} else {
			// 再使用一个0?
			res += f(1, m^1, pos-1)
		}
		for d := 1; d < b; d++ {
			res += f(1, m^(1<<d), pos-1)
		}

		return
	}

	f(0, 0, ml+1)
	return dp
}

func solve(b int, l int, r int) int {
	return calc(b, r) - calc(b, l-1)
}

func calc(b int, num int) int {
	var a []int
	for i := num; i > 0; i /= b {
		a = append(a, i%b)
	}

	var ans int

	var in, m int
	for j := len(a) - 1; j >= 0; j-- {
		for k := 0; k < a[j]; k++ {
			if k != 0 {
				in = 1
			}
			if in == 0 {
				ans += fp[b][0][m][j]
			} else {
				ans += fp[b][1][m^(1<<k)][j]
			}
		}
		if in|a[j] != 0 {
			in = 1
			m ^= (1 << a[j])
		}
	}

	ans += fp[b][in][m][0]

	return ans
}

func calc1(b int, num int) int {
	var digits []int
	for i := num; i > 0; i /= b {
		digits = append(digits, i%b)
	}
	slices.Reverse(digits)
	B := 1 << b
	dp := make([][2][2]int, B)
	dp[0][1][1] = 1

	ndp := make([][2][2]int, B)

	for _, d := range digits {
		for mask := range B {
			for eq := range 2 {
				for lz := range 2 {
					for d1 := range b {
						if dp[mask][eq][lz] == 0 {
							continue
						}
						if eq == 1 && d1 > d {
							break
						}
						if lz == 1 {
							// mask = 0
							if d1 == 0 {
								ndp[0][0][1] += dp[mask][eq][lz]
							} else {
								// d1 > 0 and d1 <= d
								neq := 1
								if d1 < d || eq == 0 {
									neq = 0
								}
								ndp[1<<d1][neq][0] += dp[mask][eq][lz]
							}
							continue
						}
						// lz == 0
						// d1 可以等于0
						new_mask := mask ^ (1 << d1)
						neq := 1
						if d1 < d || eq == 0 {
							neq = 0
						}
						ndp[new_mask][neq][0] += dp[mask][eq][lz]
					}
				}
			}
		}
		for mask := range B {
			for eq := range 2 {
				for lz := range 2 {
					dp[mask][eq][lz] = ndp[mask][eq][lz]
					ndp[mask][eq][lz] = 0
				}
			}
		}
	}

	// 不能是leading zero，mask必须是0
	return dp[0][0][0] + dp[0][1][0]
}
