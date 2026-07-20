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

type pair struct {
	first  int
	second int
}

func solve(n int, ranges [][]int, ask func(l, r int) int) int {
	cache := make(map[pair]int)

	query := func(l int, r int) int {
		if v, ok := cache[pair{l, r}]; ok {
			return v
		}
		w := ask(l+1, r+1)
		cache[pair{l, r}] = w
		return w
	}
	// when n = 3, mid = 1, when n = 4, mid = 1
	mid := (n - 1) / 2
	w := query(0, mid)

	R := make([]int, n)
	L := make([]int, n)
	for i := range R {
		R[i] = -1
		L[i] = n
	}
	for _, cur := range ranges {
		l, r := cur[0]-1, cur[1]-1
		R[l] = max(R[l], r)
		L[r] = min(L[r], l)
	}

	var ans int
	if w == 0 {
		// 那么0在后半部分
		for i := mid + 1; i < n; i++ {
			l := L[i]
			if l <= i {
				ans = max(ans, query(l, i))
			}
		}
	} else {
		// 0在前半部分
		for i := 0; i <= mid; i++ {
			r := R[i]
			if i <= r {
				ans = max(ans, query(i, r))
			}
		}
	}

	return ans
}
