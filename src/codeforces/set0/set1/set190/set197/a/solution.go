package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, r int
	fmt.Fscan(reader, &a, &b, &r)
	res := solve(a, b, r)
	fmt.Println(res)
}

func solve(a int, b int, r int) string {
	if min(a, b) < 2*r {
		return "Second"
	}

	return "First"
}
