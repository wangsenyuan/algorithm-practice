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
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	events := make([][]int, n)
	for i := range n {
		var t, u, v int
		fmt.Fscan(reader, &t, &u, &v)
		if t == 1 {
			var w int
			fmt.Fscan(reader, &w)
			events[i] = []int{t, u, v, w}
		} else {
			events[i] = []int{t, u, v}
		}
	}
	return solve(events)
}

func solve(events [][]int) []int {
	id := make(map[int]int)

	addIf := func(u int) {
		if _, ok := id[u]; !ok {
			id[u] = len(id)
		}
	}

	for _, cur := range events {
		u, v := cur[1], cur[2]

		for u != v {
			addIf(u)
			addIf(v)
			if u > v {
				u, v = v, u
			}
			// u < v
			v /= 2
		}
	}

	// 共有这么多个节点
	n := len(id)

	lca := func(u int, v int) int {
		for u != v {
			if u > v {
				u, v = v, u
			}
			v /= 2
		}
		return u
	}

	weight := make([]int, n)

	var ans []int
	for _, cur := range events {
		u, v := cur[1], cur[2]
		p := lca(u, v)
		if cur[0] == 1 {
			w := cur[3]
			for u != p {
				weight[id[u]] += w
				u /= 2
			}
			for v != p {
				weight[id[v]] += w
				v /= 2
			}
		} else {
			// u, v
			var sum int
			for u != p {
				sum += weight[id[u]]
				u /= 2
			}
			for v != p {
				sum += weight[id[v]]
				v /= 2
			}
			ans = append(ans, sum)
		}
	}

	return ans
}
