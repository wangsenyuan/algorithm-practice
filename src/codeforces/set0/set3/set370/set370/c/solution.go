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
	w, res := drive(reader)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, w)
	for _, cur := range res {
		fmt.Fprintln(writer, cur[0], cur[1])
	}
}

func drive(reader *bufio.Reader) (w int, res [][]int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	return solve(m, c)
}

type pair struct {
	first  int
	second int
}

func solve(m int, c []int) (w int, res [][]int) {
	n := len(c)

	cnt := make([]int, m+1)

	res = make([][]int, n)
	arr := make([]int, n)
	for i := range n {
		cnt[c[i]]++
		arr[i] = c[i]
		res[i] = []int{c[i], c[i]}
	}
	x := slices.Max(cnt)

	if x == n {
		return
	}

	slices.SortFunc(arr, func(x int, y int) int {
		return cmp.Or(cnt[y] - cnt[x], x - y)
	})

	buf := slices.Clone(arr)
	shift(buf, x)
	for i := range n {
		res[i][0] = arr[i]
		res[i][1] = buf[i]
		if arr[i] != buf[i] {
			w++
		}
	}
	return
}

func shift(arr []int, m int) {
	reverse(arr[:m])
	reverse(arr[m:])
	reverse(arr)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
