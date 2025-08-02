package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	return solve(n, m, k)
}

const mod = 1000000009

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

func solve(n int, m int, k int) int {
	if m < k {
		return m
	}
	// m >= k
	if n == m {
		ans := get(n, k)
		ans = add(ans, n%k)
		return ans
	}
	// m < n
	// u是k的倍数
	check := func(u int) bool {
		// u越大，那么剩余的正确答案的密度肯定越小的
		x := n - u*k
		y := m - u*k
		if y < k {
			return true
		}
		// 每隔k-1个正确答案，必须有一个错误答案
		t := y / (k - 1)
		// 还剩余a个问题
		a := y - (k-1)*t
		if a == 0 {
			// **.**.**
			b := x - (k*(t-1) + k - 1)
			return b >= 0
		}
		// a > 0
		// **.**.*
		b := x - k*t
		return a <= b
	}

	u := sort.Search(m/k, check)
	ans := get(u*k, k)
	ans = add(ans, m-u*k)
	return ans
}

func get(n int, k int) int {
	// 全部正确， 每隔k次double一次
	n = n / k
	// 第一段2 * k, 6 * k, 14 * k, 30 * k, 62 * k
	//      1, 3, 7, 15,  31
	//   (pow(2, n) - 1) * k

	res := pow(2, n)
	res = sub(res, 1)
	res = mul(res, 2)

	return mul(res, k)
}
