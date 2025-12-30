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
		var n int
		fmt.Fscan(reader, &n)
		res := solve(n)
		fmt.Fprintln(writer, res)
	}
}

func solve(n int) int {
	// gcd(n, i) = gcd(n, n - i)
	if n%2 == 1 {
		return n
	}

	return (n / 2) ^ n
}
