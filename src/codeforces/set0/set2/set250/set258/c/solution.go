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
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
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

func solve(a []int) int {
	x := slices.Max(a)

	cnt := make([]int, x+1)
	for _, v := range a {
		cnt[v]++
	}

	for i := x - 1; i > 0; i-- {
		cnt[i] += cnt[i+1]
	}
	var ans int
	for i := 1; i <= x; i++ {
		fs := factor(i)
		// 共有w个数 >= i, 它们可以使用i
		w := cnt[i]
		// 保证至少有一个数是i
		tmp := sub(pow(len(fs), w), pow(len(fs)-1, w))
		for j := len(fs) - 2; j >= 0; j-- {
			c := cnt[fs[j]] - cnt[fs[j+1]]
			// 这c个数，每个数都有(j+1)种选择
			tmp = mul(tmp, pow(j+1, c))
		}
		ans = add(ans, tmp)
	}
	return ans
}

func factor(num int) []int {
	var res []int
	for i := 1; i <= num/i; i++ {
		if num%i == 0 {
			res = append(res, i)
			if num != i*i {
				res = append(res, num/i)
			}
		}
	}
	sort.Ints(res)
	return res
}
