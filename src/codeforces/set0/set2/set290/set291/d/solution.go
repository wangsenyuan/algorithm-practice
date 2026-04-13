package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	res := solve(n, k)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, cur := range res {
		for _, x := range cur {
			fmt.Fprint(writer, x, " ")
		}
		fmt.Fprintln(writer)
	}
}

func solve(n int, k int) [][]int {
	var res [][]int
	// a[1] = n - 1
	// a[2] = n - 2
	// a[3] = n - 3
	// 1 1 1 1, 1, 0
	// 第一次所有的都变成a[i] = 2, 除了 a[n-1] = 1
	// 2 2 2 2, 1, 0
	// 第二次所有的都变成a[i] = 4, 除了 a[n-3] = 3, a[n-2] = 2, a[n-1] = 1, a[n] = 0
	// 4 4 3 2, 1, 0
	// 第三次所得都变成a[i] = 8, 除了 a[n-7] = 7, a[n-6] = 6, a[n-5] = 5
	// 5, 4
	i := n
	for d := range k {
		offset := 1 << d
		r := i - offset
		l := r - offset + 1
		c := make([]int, n+1)
		for i := 1; i < n; i++ {
			if i < l {
				c[i] = i + 1
			} else if i >= l && i < r {
				c[i] = i + offset
			} else {
				// add[0]
				c[i] = n
			}
		}
		c[n] = n
		res = append(res, c[1:])
	}
	return res
}
