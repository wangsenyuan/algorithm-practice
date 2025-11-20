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
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	var x1, x2 int
	fmt.Fscan(reader, &n, &x1, &x2)
	lines := make([][]int, n)
	for i := range n {
		var k, b int
		fmt.Fscan(reader, &k, &b)
		lines[i] = []int{k, b}
	}
	return solve(x1, x2, lines)
}

type pair struct {
	first  int
	second int
}

func solve(x1 int, x2 int, lines [][]int) bool {
	n := len(lines)

	play := func(x int) []pair {
		res := make([]pair, n)
		for i, cur := range lines {
			k, b := cur[0], cur[1]
			y := k*x + b
			res[i] = pair{y, i}
		}
		slices.SortFunc(res, func(a, b pair) int {
			return cmp.Or(a.first-b.first, a.second-b.second)
		})
		return res
	}

	w1 := play(x1)
	w2 := play(x2)

	pos := make([]int, n)
	for i, cur := range w2 {
		pos[cur.second] = i
	}

	y1 := w1[0].first
	y2 := w2[pos[w1[0].second]].first

	// 不会有两条相同的直线
	for i := 1; i < n; i++ {
		id := w1[i].second
		if w2[pos[id]].first < y2 && w1[i].first > y1 {
			return true
		}
		if w2[pos[id]].first > y2 {
			// 尽量保持y1更低
			y1 = w1[i].first
			y2 = w2[pos[id]].first
		}
	}

	return false
}
