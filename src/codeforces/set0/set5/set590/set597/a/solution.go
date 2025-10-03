package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var k, a, b int
	fmt.Fscan(reader, &k, &a, &b)
	res := solve(k, a, b)
	fmt.Println(res)
}

func solve(k int, a int, b int) int {
	if k == 1 {
		return b - a + 1
	}

	get := func(x int) int {
		if x < 0 {
			x *= -1
		}
		return x / k
	}

	if 0 < a {
		return get(b) - get(a-1)
	}
	if b < 0 {
		return get(-a) - get(-b-1)
	}
	// a <= 0 <= b
	res := get(b) + get(-a)
	res++
	return res
}
