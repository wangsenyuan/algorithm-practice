package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, change, res := drive(reader)
	fmt.Println(change)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (a []int, change int, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	change, res = solve(slices.Clone(a))
	return
}

func solve(a []int) (int, []int) {
	n := len(a)

	deg := make([]int, n)
	// 有可能有环, 和多棵树
	for i := range n {
		a[i]--
		deg[a[i]]++
	}

	res := make([]int, n)

	var que []int
	for i := range n {
		if deg[i] == 0 {
			// a leaf
			que = append(que, i)
			res[i] = a[i] + 1
		}
	}

	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		v := a[u]
		deg[v]--
		if deg[v] == 0 {
			que = append(que, v)
			res[v] = a[v] + 1
		}
	}

	var cycle [][]int

	for i := range n {
		if res[i] == 0 {
			// 是一个cycle
			j := i
			var cur []int
			for res[j] == 0 {
				cur = append(cur, j)
				k := a[j]
				res[j] = a[j] + 1
				j = k
			}
			cycle = append(cycle, cur)
		}
	}

	slices.SortFunc(cycle, func(a []int, b []int) int {
		return len(a) - len(b)
	})

	cnt := len(cycle)
	if len(cycle[0]) == 1 {
		cnt--
	}

	root := -1
	for _, cur := range cycle {
		u := cur[0]
		if root < 0 {
			res[u] = u + 1
			root = u
		} else {
			res[u] = root + 1
		}
	}

	return cnt, res
}
