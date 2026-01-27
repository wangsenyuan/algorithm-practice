package main

import (
	"bufio"
	"fmt"
	"math/bits"
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
		best, res := solve(n)
		fmt.Fprintln(writer, best)
		for _, x := range res {
			fmt.Fprint(writer, x, " ")
		}
		fmt.Fprintln(writer)
	}
}

func solve(n int) (best int, res []int) {
	if n == 1 {
		return 1, []int{1}
	}
	if n == 2 {
		return 0, []int{1, 2}
	}

	res = make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = i
	}

	h := bits.Len(uint(n))

	if n&1 == 0 {
		// even, 1 | 2 & 3 | 4
		best = 1<<h - 1
		if n&(n-1) == 0 {
			if n > 4 {
				// 4, 8, 16, ...
				// 2 | 4 & 5 | 1 & 3 |6 & 7 |8
				// a | b & (n - 1) | n
				// a | b 要等于 n - 1
				res[1], res[n-4] = res[n-4], res[1]
				res[3], res[n-3] = res[n-3], res[3]
				res = res[1:]
				return
			}
		} else {
			// 1 | 2 & 3 | 4 & 5 | 6 & 8 | 10 & 9 | 7
			res[n-2], res[n] = res[n], res[n-2]
			res[n], res[best>>1] = res[best>>1], res[n]
		}
		return best, res[1:]
	}
	// odd
	best = n
	// 2 | 4 & 5 | 6 & 7 | 1 & 3 | 8 & 9
	// a & b | (n - 1) | n
	// let c = a & b
	// 那么 c & n = (n ^ (1 << h))
	// a & b & c = c
	a := 1<<h - 1
	if a < n {
		b := n & a
		// a != b holds true
		res[a], res[n-3] = res[n-3], res[a]
		res[b], res[n-2] = res[n-2], res[b]
	} else {
		// a == n
		// 2 4 5 1 3 6 7
		res[n-3], res[1] = res[1], res[n-3]
		res[n-2], res[3] = res[3], res[n-2]
	}
	return best, res[1:]
}
