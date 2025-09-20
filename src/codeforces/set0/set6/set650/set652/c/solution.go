package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}
	foe := make([][]int, m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		foe[i] = []int{a, b}
	}
	return solve(p, foe)
}

func solve(p []int, foe [][]int) int {
	n := len(p)

	pos := make([]int, n)
	for i := range n {
		pos[p[i]-1] = i
	}

	pair := make([]int, n)
	for i := range n {
		pair[i] = -1
	}
	for _, cur := range foe {
		a, b := pos[cur[0]-1], pos[cur[1]-1]
		if a > b {
			a, b = b, a
		}
		pair[b] = max(pair[b], a)
	}

	// tr := NewSegTree(n)

	var l int

	var res int

	for r := 0; r < n; r++ {
		if pair[r] >= 0 {
			l = max(l, pair[r]+1)
		}
		res += r - l + 1
	}

	
	return res
}
