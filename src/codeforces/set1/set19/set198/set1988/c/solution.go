package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
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
		fmt.Fprintln(writer, len(res))
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func solve(n int) []int {
	var a []int

	// a[0] | a[1] = n
	// a[1] | a[2] = n

	h := bits.Len(uint(n))

	if n&(n-1) != 0 {
		for i := range h {
			if (n>>i)&1 == 1 {
				a = append(a, n^(1<<i))
			}
		}
	}

	slices.Reverse(a)

	a = append(a, n)

	return a
}
