package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	var buf bytes.Buffer
	for range tc {
		res := drive(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) int {
	var k, l1, r1, l2, r2 int
	fmt.Fscan(reader, &k, &l1, &r1, &l2, &r2)
	return solve(k, l1, r1, l2, r2)
}

func solve(k int, l1 int, r1 int, l2 int, r2 int) int {
	// x >= l1, and x <= r1
	// y >= l2  and y <= r2

	calc := func(k2 int) int {
		u := max(l1, (l2+k2-1)/k2)
		v := min(r1, r2/k2)

		if v < u {
			return 0
		}

		return v - u + 1
	}

	var res int
	k2 := 1
	for k2 <= r2 {
		res += calc(k2)
		if k2 > r2/k {
			break
		}
		k2 *= k
	}

	return res
}
