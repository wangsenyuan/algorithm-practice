package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, x int
	fmt.Fscan(reader, &n, &x)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(x, a)
}

const mod = 1000000007

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

func solve(x int, a []int) int {
	var sum int
	for _, v := range a {
		sum += v
	}

	// a[0]是最大值
	slices.Reverse(a)

	v := sum - a[0]

	var cnt int
	var i int
	for v < sum {
		for i < len(a) && sum-a[i] == v {
			i++
			cnt++
		}
		if cnt%x != 0 {
			break
		}
		cnt = cnt / x
		v++
	}

	return pow(x, v)
}
