package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var s string
	fmt.Fscan(reader, &s)
	return solve(s)
}

func solve(s string) int {
	n := len(s)

	var res int
	for i := n - 1; i >= 0; i-- {
		x := int(s[i] - '0')
		w := res % 10
		x = (x - w + 10) % 10
		res += x
	}

	return res + n
}
