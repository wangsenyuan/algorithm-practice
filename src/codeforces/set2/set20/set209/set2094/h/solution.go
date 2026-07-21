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

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	var res []int
	for range t {
		var n, q int
		fmt.Fscan(reader, &n, &q)
		a := make([]int, n)
		for i := range n {
			fmt.Fscan(reader, &a[i])
		}
		queries := make([][]int, q)
		for i := range q {
			queries[i] = make([]int, 3)
			fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
		}
		res = append(res, solve(a, queries)...)
	}
	return res
}

func solve(a []int, queries [][]int) []int {
	maxValue := 0
	for _, v := range a {
		maxValue = max(maxValue, v)
	}
	positions := make([][]int, maxValue+1)
	for i, v := range a {
		positions[v] = append(positions[v], i)
	}

	type event struct {
		pos   int
		value int
	}

	play := func(k int, l int, r int) int {
		events := make([]event, 0)
		addEvent := func(value int) {
			if value >= len(positions) {
				return
			}
			j := sort.SearchInts(positions[value], l)
			if j < len(positions[value]) && positions[value][j] <= r {
				events = append(events, event{positions[value][j], value})
			}
		}

		for d := 1; d*d <= k; d++ {
			if k%d == 0 {
				addEvent(d)
				if d*d != k {
					addEvent(k / d)
				}
			}
		}
		sort.Slice(events, func(i int, j int) bool {
			return events[i].pos < events[j].pos
		})

		var sum int
		pos := l
		for _, cur := range events {
			sum += k * (cur.pos - pos)
			for k%cur.value == 0 {
				k /= cur.value
			}
			sum += k
			pos = cur.pos + 1
		}
		sum += k * (r - pos + 1)
		return sum
	}

	ans := make([]int, len(queries))
	for i, cur := range queries {
		ans[i] = play(cur[0], cur[1]-1, cur[2]-1)
	}

	return ans
}
