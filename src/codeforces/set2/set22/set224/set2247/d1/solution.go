package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	for range q {
		var i, x int
		fmt.Fscan(reader, &i, &x)
		a[i] = x
	}
	return solve(a)
}

func solve(a []int) int {

	n := len(a)

	check := func(blockSize int) bool {
		maxVal := -1
		for i := 0; i < n; i += blockSize {
			u := a[i]
			v := a[i]
			for j := i; j < i+blockSize && j < n; j++ {
				u = min(u, a[j])
				v = max(v, a[j])
			}
			if maxVal > u {
				return false
			}
			maxVal = v
		}
		return true
	}

	if sort.IntsAreSorted(a) {
		return 0
	}

	for h := 0; 1<<h < n; h++ {
		if check(1 << (h + 1)) {
			return 1 << h
		}
	}

	return -1
}
