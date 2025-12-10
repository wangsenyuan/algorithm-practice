package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	res := solve(a)

	fmt.Println(res)
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

func sub(a, b int) int {
	return add(a, mod-b)
}

func inverse(a int) int {
	return pow(a, mod-2)
}

var i2 int

func init() {
	i2 = inverse(2)
}

func solve(a []int) int {

	f := func(state int) int {
		// 000, 101, 011, 110
		dp := make([][]int, 4)
		ndp := make([][]int, 4)
		for i := range 4 {
			dp[i] = make([]int, 3)
			ndp[i] = make([]int, 3)
		}
		// 长度为0，且当前数字所在的位置为i是的状态
		dp[0][2] = 1

		for _, v := range a {
			for l := range 4 {
				for j := range 3 {
					nl := min(l+1, 3)
					nj := (j + 1) % 3
					if v&1 == (state>>nj)&1 {
						ndp[nl][nj] = add(ndp[nl][nj], dp[l][j])
					}
				}
			}
			for l := range 4 {
				for j := range 3 {
					dp[l][j] = add(dp[l][j], ndp[l][j])
					ndp[l][j] = 0
				}
			}
		}

		var res int
		for j := range 3 {
			res = add(res, dp[3][j])
		}

		return res
	}

	res := f(0)
	res = add(res, f(3))
	res = add(res, f(5))
	res = add(res, f(6))
	return res
}

func solve1(a []int) int {

	dp := make([]int, 7)
	dp[0] = 1
	var cnt0 int
	var res int
	for _, num := range a {
		x := num & 1
		if x == 1 {
			// 011/101 都是有效状态
			res = add(res, dp[4])
			res = add(res, dp[5])
			dp[6] = add(dp[6], dp[2])
			// 状态4 =》 6， 且是一个有效状态
			dp[6] = add(dp[6], dp[4])
			// 状态5 =》 4
			dp[4] = add(dp[4], dp[5])
			dp[4] = add(dp[4], dp[1])
			dp[2] = add(dp[2], dp[0])
		} else {
			cnt0++
			//res = add(res, dp[3])
			res = add(res, dp[6])
			// 状态3 => 新的一个状态3
			//dp[3] = add(dp[3], 1)
			// 状态1 => 状态3
			dp[3] = add(dp[3], dp[1])
			// 状态0 => 状态1
			dp[1] = add(dp[1], dp[0])
			// 状态6 => 状态5
			dp[5] = add(dp[5], dp[6])
			// 状态2 => 状态5, 状态4不能变成状态5，是因为010是一个无效的组合
			dp[5] = add(dp[5], dp[2])
		}
	}
	if cnt0 >= 3 {
		tmp := pow(2, cnt0)
		// 选一个
		tmp = sub(tmp, cnt0)
		// 选两个
		tmp = sub(tmp, mul(mul(cnt0, cnt0-1), i2))
		tmp = sub(tmp, 1)

		res = add(res, tmp)
	}
	return res
}
