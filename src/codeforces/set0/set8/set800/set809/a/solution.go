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
	var n int
	fmt.Fscan(reader, &n)
	x := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &x[i])
	}
	return solve(x)
}

const mod = 1000000007

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

func solve(x []int) int {
	n := len(x)
	pw := make([]int, n+1)
	pw[0] = 1
	for i := range n {
		pw[i+1] = mul(pw[i], 2)
	}
	for i := 0; i <= n; i++ {
		pw[i] = sub(pw[i], 1)
	}
	slices.Sort(x)

	var res int
	for i := 1; i < n; i++ {
		d := x[i] - x[i-1]
		res = add(res, mul(d, mul(pw[n-i], pw[i])))
	}
	return res
}

func solve1(x []int) int {
	n := len(x)
	pw := make([]int, n+1)
	pw[0] = 1
	for i := range n {
		pw[i+1] = mul(pw[i], 2)
	}
	slices.Sort(x)
	// 如果缺定两个位置 l...r, 这段距离的贡献
	// mul(x[r] - x[l], pw[r - l - 1])
	// x[r] * pw[r - l - 1] - x[l] * pw[r-l-1]
	// x[r] * pw[r - 2] - x[1] * pw[r-2]  for l = 1
	// x[r] * pr[r - 3] - x[2] * pw[r-3]
	var res int
	//....
	var sum int
	psum := 0

	for r := 2; r <= n; r++ {
		sum = add(mul(sum, 2), x[r-2])
		psum = add(psum, pw[r-2])
		w := mul(x[r-1], psum)
		v := sub(w, sum)
		res = add(res, v)
	}

	return res
}
