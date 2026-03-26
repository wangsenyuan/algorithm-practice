package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	res := solve(n)
	fmt.Println(res)
}

func solve(n int) int {
	// 不能整除2....10
	T := 1 << 9

	var res int

	for mask := 1; mask < T; mask++ {
		p := 1
		for i := range 9 {
			if (mask>>i)&1 == 1 {
				p = lcm(p, i+2)
			}
		}
		cnt := n / p
		if bits.OnesCount(uint(mask))&1 == 1 {
			res += cnt
		} else {
			res -= cnt
		}
	}

	return n - res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	g := gcd(a, b)
	return a / g * b
}
