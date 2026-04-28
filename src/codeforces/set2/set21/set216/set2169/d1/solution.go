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
		var x, y, k int
		fmt.Fscan(reader, &x, &y, &k)
		res := solve(x, y, k)
		fmt.Fprintln(writer, res)
	}
}

func solve(x int, y int, k int) int {
	check := func(w int) bool {
		// 1 .... w 还剩下多少
		m := w
		for range x {
			m -= m / y
			if m < k {
				return false
			}
		}
		return true
	}

	res := sort.Search(1e12+1, check)
	if res > 1e12 {
		return -1
	}
	return res
}
