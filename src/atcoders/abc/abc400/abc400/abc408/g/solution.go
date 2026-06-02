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
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var a, b, c, d int
	fmt.Fscan(reader, &a, &b, &c, &d)
	return solve(a, b, c, d)
}

func solve(a int, b int, c int, d int) int {
	var f func(a int, b int, c int, d int) (p int, q int)
	f = func(a int, b int, c int, d int) (p int, q int) {
		if a < b {
			if c > d {
				return 1, 1
			}
			// c <= d
			p, q = f(d, c, b, a)
			return q, p
		}
		n := a / b
		p, q = f(a-n*b, b, c-n*d, d)
		p += n * q
		return
	}
	_, q := f(a, b, c, d)
	return q
}
