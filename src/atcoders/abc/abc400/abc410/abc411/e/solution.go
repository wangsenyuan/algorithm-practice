package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
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

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
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

func inv(a int) int {
	return pow(a, mod-2)
}

func div(a, b int) int {
	return mul(a, inv(b))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, 6)
		for j := range 6 {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a)
}

func solve(a [][]int) int {
	var s []int
	for _, vs := range a {
		for _, x := range vs {
			s = append(s, x)
		}
	}
	slices.Sort(s)
	s = slices.Compact(s)
	k := len(s)
	upd := make([][]int, k)
	for j, vs := range a {
		for _, x := range vs {
			i := sort.SearchInts(s, x)
			upd[i] = append(upd[i], j)
		}
	}
	b := make([]int, len(a))
	var ans int
	prod := 1
	zero := len(a)

	for i := range k - 1 {
		for _, j := range upd[i] {
			if b[j] == 0 {
				zero--
			} else {
				prod = div(prod, b[j])
			}
			b[j]++
			prod = mul(prod, b[j])
		}

		if zero == 0 {
			ans = sub(ans, mul(prod, s[i+1]-s[i]))
		}
	}

	ans = div(ans, pow(6, len(a)))

	ans = add(ans, s[k-1])

	return ans
}
