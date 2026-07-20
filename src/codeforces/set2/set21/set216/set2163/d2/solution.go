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

	var t int
	fmt.Fscan(reader, &t)
	for range t {
		var n, q int
		fmt.Fscan(reader, &n, &q)
		ranges := make([][]int, q)
		for i := range q {
			ranges[i] = make([]int, 2)
			fmt.Fscan(reader, &ranges[i][0], &ranges[i][1])
		}

		ask := func(l, r int) int {
			fmt.Fprintf(writer, "? %d %d\n", l, r)
			writer.Flush()
			var res int
			fmt.Fscan(reader, &res)
			return res
		}

		ans := solve(n, ranges, ask)
		fmt.Fprintf(writer, "! %d\n", ans)
		writer.Flush()
	}
}

func solve(n int, ranges [][]int, ask func(l, r int) int) int {
	L := make([]int, n+1)
	R := make([]int, n+1)
	for i := range n + 1 {
		L[i] = n + 1
		R[i] = -1
	}

	for _, cur := range ranges {
		l, r := cur[0], cur[1]
		L[r] = min(L[r], l)
		R[l] = max(R[l], r)
	}

	if R[1] == n {
		return n
	}

	var todo [][]int
	var next int
	for i := 1; i <= n; i++ {
		if R[i] > -1 && R[i] > next && L[R[i]] == i {
			todo = append(todo, []int{i, R[i]})
			next = R[i]
		}
	}

	// mex(l...r) = min(mex(l...n), mex(1...r))
	// 如果 mex(l...n) < mex(1...r), 那么说明有一个关键的数出现在区间 [1...l-1]
	// 那么就应该往左边寻找

	l, r := 0, len(todo)
	var ans int

	for l < r {
		mid := (l + r) / 2
		a := ask(1, todo[mid][1])
		b := ask(todo[mid][0], n)
		ans = max(ans, min(a, b))
		if a < b {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return ans
}
