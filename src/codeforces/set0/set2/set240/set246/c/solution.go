package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, cur := range res {
		fmt.Fprintf(writer, "%d ", len(cur))
		for _, v := range cur {
			fmt.Fprintf(writer, "%d ", v)
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) (k int, a []int, res [][]int) {
	var n int
	fmt.Fscan(reader, &n, &k)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(k, a)
	return
}

func solve(k int, a []int) (res [][]int) {
	n := len(a)

	slices.Sort(a)
	slices.Reverse(a)

	for i := range n {
		res = append(res, []int{a[i]})
		if len(res) == k {
			break
		}
	}
	var pref []int

	for i := 0; len(res) < k; i++ {
		pref = append(pref, a[i])
		for j := i + 1; j < n && len(res) < k; j++ {
			tmp := slices.Clone(pref)
			tmp = append(tmp, a[j])
			res = append(res, tmp)
		}
	}

	return res
}
