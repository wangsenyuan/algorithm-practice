package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	res := solve(n)
	fmt.Println(res[0], res[1])
}

const inf = 1 << 60

func solve(n int) []int {
	// 当 a = 2时， b = 3时，c = n
	mx := max(2*3*(n+2), 3*3*(n+1)) - n

	mn := mx

	check := func(a int) int {
		area := n / (a - 1)
		// (b - 2) * (c - 2) = area
		// b和c尽量相等
		b2 := int(math.Sqrt(float64(area)))

		for x := b2 + 1; x > 0; x-- {
			if area%x == 0 {
				y := area / x
				// x * y = area
				return a * (x + 2) * (y + 2)
			}
		}

		return inf
	}

	for a := 1; a <= n/a; a++ {
		// (a - 1) * (b - 2) * (c - 2) = n
		// 且 b和c足够接近
		if n%a == 0 {
			tmp := check(a + 1)
			mn = min(mn, tmp-n)
			if n != a*a {
				tmp = check(n/a + 1)
				mn = min(mn, tmp-n)
			}
		}
	}

	return []int{mn, mx}
}
