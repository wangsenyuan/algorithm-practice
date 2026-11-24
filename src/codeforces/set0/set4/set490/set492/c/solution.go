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

	fmt.Fprintln(writer, drive(reader))
}

func drive(reader *bufio.Reader) int64 {
	var n int
	var r, avg int64
	fmt.Fscan(reader, &n, &r, &avg)

	a := make([]int64, n)
	b := make([]int64, n)
	for i := range n {
		fmt.Fscan(reader, &a[i], &b[i])
	}

	return solve(n, r, avg, a, b)
}

func solve(n int, r, avg int64, a, b []int64) int64 {
	type exam struct {
		grade int64
		cost  int64
	}

	exams := make([]exam, n)
	var sum int64
	for i := range n {
		exams[i] = exam{grade: a[i], cost: b[i]}
		sum += a[i]
	}

	sort.Slice(exams, func(i, j int) bool {
		return exams[i].cost < exams[j].cost
	})

	want := avg * int64(n)

	var res int64

	for _, cur := range exams {
		if want <= sum {
			break
		}
		w := min(r-cur.grade, want-sum)
		res += w * cur.cost
		sum += w
	}

	if sum < want {
		return -1
	}
	return res
}
