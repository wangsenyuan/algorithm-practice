package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

type trampoline struct {
	x, y, p int
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	pts := make([]trampoline, n)
	for i := range n {
		fmt.Fscan(reader, &pts[i].x, &pts[i].y, &pts[i].p)
	}
	return solve(pts)
}

func solve(pts []trampoline) int {
	// TODO: solve by hand first.
	n := len(pts)
	vis := make([]bool, n)
	que := make([]int, n)

	canReach := func(u int, v int, s int) bool {
		dx := abs(pts[u].x - pts[v].x)
		dy := abs(pts[u].y - pts[v].y)
		return pts[u].p*s >= dx+dy
	}

	play := func(s int, u1 int) bool {
		clear(vis)
		vis[u1] = true
		var head, tail int
		que[head] = u1
		head++
		for tail < head {
			u := que[tail]
			tail++
			for v := range n {
				if !vis[v] && canReach(u, v, s) {
					vis[v] = true
					que[head] = v
					head++
				}
			}
		}

		return head == n
	}

	const inf = 4e9 + 10

	ans := sort.Search(inf, func(s int) bool {
		for u := range n {
			if play(s, u) {
				return true
			}
		}
		return false
	})

	return ans
}

func abs(a int) int {
	return max(a, -a)
}
