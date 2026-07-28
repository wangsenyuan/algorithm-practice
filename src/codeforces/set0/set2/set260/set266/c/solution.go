package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	ops, _, _ := drive(reader)
	fmt.Fprintln(writer, len(ops))
	for _, op := range ops {
		fmt.Fprintln(writer, op[0], op[1], op[2])
	}
}

func drive(reader *bufio.Reader) (res [][]int, n int, ones [][]int) {
	fmt.Fscan(reader, &n)
	ones = make([][]int, n-1)
	buf := make([][]int, n-1)
	for i := range n - 1 {
		ones[i] = make([]int, 2)
		fmt.Fscan(reader, &ones[i][0], &ones[i][1])
		buf[i] = slices.Clone(ones[i])
	}
	res = solve(n, buf)
	return
}

func solve(n int, ones [][]int) [][]int {
	var res [][]int

	slices.SortFunc(ones, func(a []int, b []int) int {
		return cmp.Or(b[0]-a[0], a[1]-b[1])
	})

	var f func(r int, ones [][]int)
	col := make([]int, n+1)

	f = func(n int, ones [][]int) {
		if len(ones) == 0 {
			return
		}
		clear(col[:n+1])
		// len(ones) < r
		// 如果r列存在1, 那么找到那个没有1的列, 把它给swap过来
		for _, cur := range ones {
			r, c := cur[0], cur[1]
			col[c] = r
		}
		if col[n] != 0 {
			c1 := 1
			for col[c1] != 0 {
				c1++
			}
			res = append(res, []int{2, c1, n})
			// 所以那些列是r的需要交换到c1列去
			for _, cur := range ones {
				if cur[1] == n {
					cur[1] = c1
				}
			}
		}
		// col[r] == 0
		var i int
		for i < len(ones) && ones[i][0] == n {
			i++
		}

		if i == 0 {
			r := ones[0][0]
			res = append(res, []int{1, r, n})
			for i < len(ones) && ones[i][0] == r {
				i++
			}
		}

		f(n-1, ones[i:])
	}

	f(n, ones)

	return res
}
