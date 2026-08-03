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
	var r, h int
	fmt.Fscan(reader, &r, &h)
	return solve(r, h)
}

func solve(r int, h int) int {
	res := h / r * 2
	x := h % r
	if 2*x < r {
		res++
	} else if 4*int64(x)*int64(x) < 3*int64(r)*int64(r) {
		res += 2
	} else {
		res += 3
	}
	return res
}
