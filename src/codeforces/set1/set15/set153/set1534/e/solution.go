package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const inf = 0x3f3f3f3f

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	dis := make([]int, 2001)
	par := make([]int, 2001)
	for i := range dis {
		dis[i] = inf
		par[i] = -1
	}

	dis[0] = 0
	que := []int{0}
	for len(que) > 0 {
		c := que[0]
		que = que[1:]
		for i := 0; i <= k; i++ {
			if i <= n-c && k-i <= c {
				to := c + i - (k - i)
				if dis[to] == inf {
					dis[to] = dis[c] + 1
					par[to] = c
					que = append(que, to)
				}
			}
		}
	}

	if dis[n] == inf {
		fmt.Fprintln(writer, -1)
		return
	}

	path := make([]int, 0)
	for c := n; c != -1; c = par[c] {
		path = append(path, c)
	}
	slices.Reverse(path)

	nosel := make([]int, n)
	for i := range n {
		nosel[i] = i + 1
	}
	var sel []int

	ans := 0
	for i := 0; i+1 < len(path); i++ {
		a, b := path[i], path[i+1]
		d := b - a
		nsel := (d + k) / 2
		nnosel := k - nsel

		csel := make([]int, nsel)
		copy(csel, nosel[len(nosel)-nsel:])
		nosel = nosel[:len(nosel)-nsel]

		cnosel := make([]int, nnosel)
		copy(cnosel, sel[len(sel)-nnosel:])
		sel = sel[:len(sel)-nnosel]

		query := append(csel, cnosel...)
		fmt.Fprint(writer, "? ")
		for j, x := range query {
			if j > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, x)
		}
		fmt.Fprintln(writer)
		writer.Flush()

		nosel = append(nosel, cnosel...)
		sel = append(sel, csel...)

		var res int
		fmt.Fscan(reader, &res)
		ans ^= res
	}

	fmt.Fprintf(writer, "! %d\n", ans)
}
