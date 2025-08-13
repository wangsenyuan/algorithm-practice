package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, n int
	fmt.Scanf("%d %d %d", &x, &y, &n)
	res := solve(x, y, n)
	fmt.Printf("%d/%d\n", res[0], res[1])
}

func solve(x int, y int, n int) []int {
	g := gcd(x, y)
	x /= g
	y /= g
	if y <= n {
		return []int{x, y}
	}
	res := []int{0, 1}
	best := float64(x) / float64(y)

	update := func(a int, b int) {
		if gcd(a, b) > 1 {
			return
		}
		dist := (float64(a*y) - float64(x*b)) / float64(b*y)
		dist = math.Abs(dist)
		if dist < best {
			best = dist
			res = []int{a, b}
		}
	}

	for b := 1; b <= n; b++ {
		a := x * b / y
		update(a, b)
		a = (x*b + y - 1) / y
		update(a, b)
	}

	return res
}

func gcd(a int, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
