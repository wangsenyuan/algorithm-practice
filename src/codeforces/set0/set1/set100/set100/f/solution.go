package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) string {
	n := len(a)
	res := make([]int, n+1)

	for mask := range 1 << n {
		var cnt int
		prod := 1
		for i := range n {
			if (mask>>i)&1 == 1 {
				cnt++
				prod *= a[i]
			}
		}
		res[cnt] += prod
	}

	if res[n] < 0 {
		for i := range n + 1 {
			res[i] *= -1
		}
	}

	var buf bytes.Buffer
	for i := n; i >= 0; i-- {
		coeff := res[i]
		if coeff == 0 {
			continue
		}
		if coeff > 0 {
			buf.WriteByte('+')
		} else {
			buf.WriteByte('-')
			coeff = -coeff
		}
		if coeff > 1 {
			fmt.Fprintf(&buf, "%d*", coeff)
		}
		if i > 0 {
			buf.WriteByte('X')
			if i > 1 {
				fmt.Fprintf(&buf, "^%d", i)
			}
		} else {
			fmt.Fprintf(&buf, "%d", coeff)
		}
	}
	ans := buf.String()

	return ans[1:]
}
