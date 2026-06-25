package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		_, res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) (a []int, p []int) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	p = solve(a)
	return
}

func solve(a []int) []int {
	n := len(a)
	type state struct {
		sum int
		pos int
	}
	arr := make([]state, n)
	var sum int
	for i := range n {
		arr[i] = state{sum, i}
		sum += a[i]
	}
	slices.SortFunc(arr, func(a state, b state) int {
		return cmp.Or(b.sum-a.sum, a.pos-b.pos)
	})
	res := make([]int, n)
	for i, cur := range arr {
		res[cur.pos] = i + 1
	}
	return res
}
