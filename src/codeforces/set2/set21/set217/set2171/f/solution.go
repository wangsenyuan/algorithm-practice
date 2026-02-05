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
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		_, res := drive(reader)
		if len(res) == 0 {
			fmt.Fprintln(writer, "No")
			continue
		}

		fmt.Fprintln(writer, "Yes")
		for _, e := range res {
			fmt.Fprintln(writer, e[0], e[1])
		}
	}
}

func drive(reader *bufio.Reader) (p []int, res [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	p = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	res = solve(p)
	return
}

func solve(p []int) [][]int {
	// n := len(p)

	var active []int

	var res [][]int

	for _, v := range p {
		fa := v
		for len(active) > 0 && active[len(active)-1] < v {
			u := active[len(active)-1]
			active = active[:len(active)-1]
			fa = min(fa, u)
			res = append(res, []int{u, v})
		}
		// 把最小的放进去
		active = append(active, fa)
	}

	if len(active) != 1 {
		return nil
	}

	return res
}
