package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	x := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &x[i])
	}
	return solve(x)
}

const inf = 1 << 60

func solve(x []int) int {
	n := len(x)
	if n == 1 {
		return 0
	}
	var res int

	j := 1

	var sum int
	lo, hi := 0, inf
	for i := 1; i < n; i++ {
		sum = (x[i] - x[i-1]) - sum

		if i&1 == j&1 {
			hi = min(hi, sum)
			if i+1 < n {
				lo = max(lo, sum-(x[i+1]-x[i]))
			}
		} else {
			lo = max(lo, -sum)
			if i+1 < n {
				hi = min(hi, (x[i+1]-x[i])-sum)
			}
		}
		if lo >= hi {
			res--
			lo = 0
			hi = inf
			sum = 0
			j = i + 1
		}
	}

	return res + n - 1
}
