package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Println(drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var c int
	var d int
	fmt.Fscan(reader, &c, &d)
	return solve(c, d)
}

func isqrt(x int) int {
	rt := int(math.Sqrt(float64(x)))
	if rt*rt > x {
		rt--
	}
	return rt
}

func solve(c int, d int) int {
	ans := 0
	for p10 := 1; p10 <= c+d; p10 *= 10 {
		l := c*p10*10 + max(c+1, p10)
		r := c*p10*10 + min(c+d, p10*10-1)
		if l <= r {
			ans += isqrt(r) - isqrt(l-1)
		}
	}

	return ans
}
