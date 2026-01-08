package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%d/%d\n", res[0], res[1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const mod = 1e9 + 7

func add(a int, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a int, b int) int {
	return add(a, mod-b)
}

func mul(a int, b int) int {
	return a * b % mod
}

func pow(a int, b int) int {
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

func inverse(a int) int {
	return pow(a, mod-2)
}

func div(a int, b int) int {
	return mul(a, inverse(b))
}

func solve(a []int) []int {

	// 需要知道n的奇偶性
	// 只有所有的都为1的时候，才是奇数
	// 只要有一个偶数，n就是偶数
	odd := true
	for _, v := range a {
		if v&1 == 0 {
			odd = false
			break
		}
	}
	// 偶数
	res := 2
	// 要计算2 ^^ (n - 1)
	for _, v := range a {
		res = pow(res, v)
	}

	// res = pow(2, n - 1)
	res = div(res, 2)

	var p int
	if odd {
		p = div(sub(res, 1), 3)
	} else {
		p = div(add(res, 1), 3)
	}

	return []int{p, res}
}
