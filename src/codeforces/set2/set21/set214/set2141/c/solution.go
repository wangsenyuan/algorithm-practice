package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	res := solve(n)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, len(res))
	for _, cur := range res {
		fmt.Fprintln(writer, cur)
	}
}

func solve(n int) []string {

	pushBack := func(i int) string {
		return fmt.Sprintf("pushback a[%d]", i-1)
	}

	var res []string
	d := 1
	for l := 1; l <= n; l++ {
		if d == 1 {
			r := l
			for r <= n {
				res = append(res, pushBack(r))
				res = append(res, "min")
				r++
			}
			// r > n
			d = -1
		} else {
			r := n
			// d = -1
			// r = n
			for r > l {
				res = append(res, "min")
				res = append(res, "popback")
				r--
			}
			// l = r 但是还没有查询
			// res = append(res, "min")
			res = append(res, "min")
			d = 1
		}

		if l < n {
			res = append(res, "popfront")
		}
	}

	return res
}
