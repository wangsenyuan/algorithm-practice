package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := drive(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) [][]int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	segs := make([][]int, n)
	for i := 0; i < n; i++ {
		segs[i] = make([]int, 2)
		fmt.Fscan(reader, &segs[i][0], &segs[i][1])
	}
	return solve(k, segs)
}

type event struct {
	pos int
	tp  int
}

func solve(k int, segs [][]int) [][]int {
	n := len(segs)
	events := make([]event, 2*n)
	for i := range n {
		l, r := segs[i][0], segs[i][1]
		events[2*i] = event{l, -1}
		events[2*i+1] = event{r, 1}
	}

	slices.SortFunc(events, func(a, b event) int {
		return cmp.Or(a.pos-b.pos, a.tp-b.tp)
	})

	var res [][]int
	var l = -inf
	var level int
	for _, cur := range events {
		level -= cur.tp
		switch level {
		case k:
			if cur.tp == -1 {
				l = cur.pos
			}
		case k - 1:
			if cur.tp == 1 {
				res = append(res, []int{l, cur.pos})
			}
		}
	}
	return res
}

const inf = 1 << 60
