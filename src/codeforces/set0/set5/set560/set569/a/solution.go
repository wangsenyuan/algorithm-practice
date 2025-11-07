package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t, s, q int
	fmt.Fscan(reader, &t, &s, &q)
	res := solve(t, s, q)
	fmt.Println(res)
}

func solve(t int, s int, q int) int {
	var res int
	for s < t {
		res++
		if s > t/q {
			break
		}
		s *= q
	}
	return res
}
