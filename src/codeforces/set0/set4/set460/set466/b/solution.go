package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, a1, b1 := drive(reader)
	fmt.Println(s)
	fmt.Println(a1, b1)
}

func drive(reader *bufio.Reader) (int, int, int) {
	var n, a, b int
	fmt.Fscan(reader, &n, &a, &b)
	res := solve(n, a, b)
	return res[0], res[1], res[2]
}

func solve(n int, a int, b int) []int {
	T := 6 * n
	if a*b >= T {
		return []int{a * b, a, b}
	}
	sqrtT := int(math.Sqrt(float64(T))) + 2

	bestArea := math.MaxInt64
	bestA, bestB := a, b

	update := func(a1, b1 int) {
		if area := a1 * b1; area < bestArea {
			bestArea, bestA, bestB = area, a1, b1
		}
	}

	// Iterate a1 from a up to sqrt(T); for each, use the smallest valid b1
	for a1 := a; a1 <= sqrtT; a1++ {
		b1 := (T + a1 - 1) / a1
		if b1 < b {
			b1 = b
		}
		update(a1, b1)
	}

	// Iterate b1 from b up to sqrt(T); for each, use the smallest valid a1
	for b1 := b; b1 <= sqrtT; b1++ {
		a1 := (T + b1 - 1) / b1
		if a1 < a {
			a1 = a
		}
		update(a1, b1)
	}

	return []int{bestArea, bestA, bestB}
}
