package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	roads := make([][]int, n)
	for i := range n {
		roads[i] = make([]int, 3)
		fmt.Fscan(reader, &roads[i][0], &roads[i][1], &roads[i][2])
	}
	return solve(n, roads)
}

func solve(n int, roads [][]int) int {

	g := make([][]int, n)
	var sum int
	for i, cur := range roads {
		u, v := cur[0]-1, cur[1]-1
		g[u] = append(g[u], i)
		g[v] = append(g[v], i)
		sum += cur[2]
	}

	s := 0
	eid := g[0][0]
	var sum2 int

	for {
		u, v, w := roads[eid][0]-1, roads[eid][1]-1, roads[eid][2]
		// u -> v
		var t int
		if v == s {
			sum2 += w
			t = u
		} else {
			t = v
		}

		if t == 0 {
			break
		}

		// len(g[t]) == 2
		for _, i := range g[t] {
			if eid != i {
				eid = i
				break
			}
		}

		s = t
	}

	return min(sum2, sum-sum2)
}
