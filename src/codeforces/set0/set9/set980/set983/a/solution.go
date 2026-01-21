package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var p, q, b int
		fmt.Fscan(reader, &p, &q, &b)
		res := solve(p, q, b)
		fmt.Fprintln(writer, res)
	}
}

func solve(p int, q int, b int) string {
	q /= gcd(p, q)
	b = gcd(q, b)

	for b != 1 {
		for q%b == 0 {
			q /= b
		}
		b = gcd(q, b)
	}
	if q == 1 {
		return "Finite"
	}
	return "Infinite"
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
