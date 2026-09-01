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

	for _, ans := range drive(reader) {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	friends := make([][]int, m)
	for i := range m {
		friends[i] = make([]int, 2)
		fmt.Fscan(reader, &friends[i][0], &friends[i][1])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := range q {
		var typ int
		fmt.Fscan(reader, &typ)
		if typ == 3 {
			queries[i] = []int{3}
			continue
		}
		var u, v int
		fmt.Fscan(reader, &u, &v)
		queries[i] = []int{typ, u, v}
	}
	return solve(n, friends, queries)
}

func solve(n int, friends [][]int, queries [][]int) []int {
	// stronger[i] = number of friends with higher power. A noble survives
	// the killing process iff this count is 0: anyone with a stronger
	// neighbor eventually dies, and anyone without one never becomes
	// vulnerable (they only lose weaker friends, then sit isolated).
	stronger := make([]int, n+1)
	alive := n

	add := func(u, v int) {
		if u > v {
			u, v = v, u
		}
		if stronger[u] == 0 {
			alive--
		}
		stronger[u]++
	}
	remove := func(u, v int) {
		if u > v {
			u, v = v, u
		}
		stronger[u]--
		if stronger[u] == 0 {
			alive++
		}
	}

	for _, e := range friends {
		add(e[0], e[1])
	}

	var ans []int
	for _, q := range queries {
		switch q[0] {
		case 1:
			add(q[1], q[2])
		case 2:
			remove(q[1], q[2])
		default:
			ans = append(ans, alive)
		}
	}
	return ans
}
