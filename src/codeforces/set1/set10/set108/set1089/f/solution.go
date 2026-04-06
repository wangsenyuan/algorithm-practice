package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int64
	fmt.Fscan(in, &n)

	ok, ans := solve(n)
	if !ok {
		fmt.Println("NO")
		return
	}

	fmt.Println("YES")
	fmt.Println(len(ans))
	for _, cur := range ans {
		fmt.Println(cur[0], cur[1])
	}
}

func solve(n int64) (bool, [][2]int64) {
	x := primePowerPart(n)
	y := n / x
	if x == 1 || y == 1 {
		return false, nil
	}

	_, inv, _ := exgcd(y, x)
	inv %= x
	if inv < 0 {
		inv += x
	}

	a := (x - inv) % x
	if a == 0 {
		a = x - 1
	}
	b := (n - 1 - a*y) / x

	if !(1 <= a && a < x && 1 <= b && b < y) {
		return false, nil
	}

	return true, [][2]int64{{a, x}, {b, y}}
}

func primePowerPart(n int64) int64 {
	for p := int64(2); p*p <= n; p++ {
		if n%p == 0 {
			res := int64(1)
			for n%p == 0 {
				n /= p
				res *= p
			}
			return res
		}
	}
	return n
}

func exgcd(a, b int64) (g, x, y int64) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := exgcd(b, a%b)
	return g, y1, x1-(a/b)*y1
}
