package main

import (
	"bufio"
	"fmt"
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

const N = 5010

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(nums ...int) int {
	res := 1
	for _, num := range nums {
		res = res * num % mod
	}
	return res
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func inverse(a int) int {
	return pow(a, mod-2)
}

var F [N]int
var invF [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(F[i-1], i)
	}
	invF[N-1] = inverse(F[N-1])
	for i := N - 1; i > 0; i-- {
		invF[i-1] = mul(invF[i], i)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n2 := len(a)
	n := n2 / 2

	last := -1
	var arr []int
	for _, v := range a {
		if last != v {
			arr = append(arr, 1)
		} else {
			arr[len(arr)-1]++
		}
		last = v
	}
	m := len(arr)

	dp := make([]int, n+1)
	dp[0] = 1

	get := func(j int) int {
		if j == n {
			return j
		}
		return max(0, j-1)
	}

	perm := func(x int, k int) int {
		if x < 0 || k < 0 || x < k {
			return 0
		}
		return mul(F[x], invF[x-k])
	}

	// ndp := make([]int, n+1)

	var suf int
	for i := m - 1; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			free0 := get(j) - (suf - j)
			cur := mul(dp[j], perm(free0, arr[i]))

			if j > 0 {
				free1 := get(j-1) - (suf - (j - 1))
				cur = add(cur, mul(arr[i], dp[j-1], perm(free1, arr[i]-1)))
			}

			dp[j] = cur
		}
		suf += arr[i]
	}

	res := dp[n]
	for _, c := range arr {
		res = mul(res, invF[c])
	}
	return res
}
