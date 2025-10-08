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
	var p int
	fmt.Fscan(reader, &p)
	a := make([]int, 2)
	x := make([]int, 3)
	y := make([]int, 3)
	fmt.Fscan(reader, &a[0], &a[1])
	fmt.Fscan(reader, &x[0], &y[0])
	b := make([]int, 2)
	fmt.Fscan(reader, &b[0], &b[1])
	fmt.Fscan(reader, &x[1], &y[1])
	return solve(p, a, b, x, y)
}

func solve(p int, a []int, b []int, x []int, y []int) int {
	next := func(c int, w int) int {
		return (c*x[w] + y[w]) % p
	}
	var ans int
	for a[0] != a[1] && ans < p+20 {
		ans++
		a[0] = next(a[0], 0)
		b[0] = next(b[0], 1)
	}
	if a[0] != a[1] {
		return -1
	}
	if b[0] == b[1] {
		return ans
	}
	var o int
	cur := a[0]
	x[2] = x[1]
	y[2] = y[1]
	x[1] = 1
	y[1] = 0
	for {
		cur = next(cur, 0)
		o++
		x[1] = x[1] * x[2] % p
		y[1] = y[1] * x[2] % p
		y[1] = (y[1] + y[2]) % p
		if o >= p+20 || cur == a[1] {
			break
		}
	}
	if cur != a[1] {
		return -1
	}
	O := 0
	cur = b[0]
	for {
		cur = next(cur, 1)
		O++
		if O >= p+20 || cur == b[1] {
			break
		}
	}
	if cur != b[1] {
		return -1
	}
	ans += o * O
	return ans
}
