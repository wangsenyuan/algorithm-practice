package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for i, x := range res {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, x)
	}
	fmt.Fprintln(writer)
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	ops := make([][]int, q)
	for i := range q {
		ops[i] = make([]int, 2)
		fmt.Fscan(reader, &ops[i][0], &ops[i][1])
	}
	return solve(n, ops)
}

func solve(n int, ops [][]int) []int {

	diff0 := make([]int, n+2)
	diff1 := make([]int, n+2)

	for _, op := range ops {
		x, y := op[0], op[1]
		// a[j] += max(y-|x-j|, 0)，正贡献在 [x-y+1, x+y-1]
		// [l, x] 上是 j+(y-x)，斜率 +1
		// [x+1, r] 上是 -j+(x+y)，斜率 -1
		l := max(1, x-y+1)
		r := min(n, x+y-1)

		diff1[l]++
		diff1[x+1]--
		diff0[l] += l - 1 + y - x
		diff0[x+1] -= y
		if x < r {
			diff1[x+1]--
			diff1[r+1]++
			diff0[x+1] += y
			diff0[r+1] += r - x - y
		}
	}

	a := make([]int, n+1)

	var diff int
	for i := 1; i <= n; i++ {
		diff0[i] += diff0[i-1]
		diff1[i] += diff1[i-1]
		diff += diff1[i]
		a[i] = diff0[i] + diff
	}

	return a[1:]
}
