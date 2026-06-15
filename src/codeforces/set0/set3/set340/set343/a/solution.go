package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b int
	fmt.Fscan(reader, &a, &b)
	fmt.Println(solve(a, b))
}

func solve(a int, b int) int {
	res := 0
	for b > 0 {
		res += a / b
		a, b = b, a%b
	}
	return res
}
