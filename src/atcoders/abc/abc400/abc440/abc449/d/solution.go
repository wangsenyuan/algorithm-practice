package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var L, R, D, U int
	fmt.Fscan(reader, &L, &R, &D, &U)
	res := solve(L, R, D, U)
	fmt.Println(res)
}

func solve(L int, R int, D int, U int) int {
	if R < 0 {
		// 左右翻转
		return solve(-R, -L, D, U)
	}
	// R >= 0. 但是 L有可能小于0
	if U < 0 {
		// 上下翻转
		return solve(L, R, -U, -D)
	}

	if L < 0 {
		return solve(1, R, D, U) + solve(0, -L, D, U)
	}
	if D < 0 {
		return solve(L, R, 1, U) + solve(L, R, 0, -D)
	}

	return help1(L, R, D, U)
}

func help1(l int, r int, d int, u int) int {
	if l > r || d > u {
		return 0
	}
	// l >= 0, r >= 0, d >= 0, u >= 0
	res := help2(r, u)
	if l > 0 {
		res -= help2(l-1, u)
	}
	if d > 0 {
		res -= help2(r, d-1)
	}
	if l > 0 && d > 0 {
		res += help2(l-1, d-1)
	}
	return res
}

func help2(x int, y int) int {
	if x < 0 || y < 0 {
		return 0
	}
	if y > x {
		y, x = x, y
	}
	k := y/2 + 1
	return k*k*2 - k + (x/2-y/2)*(y+1)
}
