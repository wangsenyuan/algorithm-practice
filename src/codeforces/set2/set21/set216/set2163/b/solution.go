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
		_, _, ok, res := drive(reader)
		if !ok {
			fmt.Fprintln(writer, "-1")
		} else {
			fmt.Fprintln(writer, len(res))
			for _, e := range res {
				fmt.Fprintln(writer, e[0], e[1])
			}
		}
	}
}

func drive(reader *bufio.Reader) (p []int, x string, ok bool, res [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	p = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	fmt.Fscan(reader, &x)
	ok, res = solve(p, x)
	return
}

func solve(p []int, x string) (ok bool, res [][]int) {
	n := len(p)
	pos := make([]int, n)
	for i, v := range p {
		pos[v-1] = i
	}

	if x[pos[0]] == '1' || x[pos[n-1]] == '1' || x[0] == '1' || x[n-1] == '1' {
		return false, nil
	}

	add := func(l int, r int) {
		res = append(res, []int{l + 1, r + 1})
	}
	l, r := min(pos[0], pos[n-1]), max(pos[0], pos[n-1])
	add(l, r)
	add(0, pos[0])
	add(0, pos[n-1])
	add(pos[0], n-1)
	add(pos[n-1], n-1)
	return true, res
}
