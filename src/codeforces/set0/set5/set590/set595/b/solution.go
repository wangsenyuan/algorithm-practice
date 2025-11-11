package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n/k)
	for i := range n / k {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n/k)
	for i := range n / k {
		fmt.Fscan(reader, &b[i])
	}
	return solve(n, k, a, b)
}

const mod = 1e9 + 7

func mul(a, b int) int {
	return a * b % mod
}

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

func solve(n int, k int, a []int, b []int) int {
	w := 1
	for range k {
		w *= 10
	}
	// w是上限

	calc := func(down int, up int, fa int) int {
		res := up / fa
		if up >= 0 {
			res++
		}
		res -= max(0, (down-1)/fa)
		if down-1 >= 0 {
			res--
		}
		return res
	}

	ans := 1
	for i := range len(a) {
		x := calc(0, w-1, a[i])
		y := calc(b[i]*(w/10), (b[i]+1)*(w/10)-1, a[i])
		ans = mul(ans, sub(x, y))
	}

	return ans
}
